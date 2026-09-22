import { useEffect, useRef, useState } from 'react';
import { Button, Card, CardBody, CardFooter, CardHeader, Chip, Progress } from '@heroui/react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import { CheckCircle, ExternalLink, Network, RefreshCw, Square } from 'lucide-react';
import { RoutesLocation } from '@/components/routes';
import { ErrorModal } from '@/components/pages/auth/error';
import { useMutationSessionEnsure } from '@/helpers/queries/session/use-mutation-session-ensure';
import { useQuerySessionGet } from '@/helpers/queries/session/use-query-session-get';
import { useMutationSessionStop } from '@/helpers/queries/session/use-mutation-session-stop';
import { useMutationSessionCheck } from '@/helpers/queries/session/use-mutation-session-check';
import { useMutationSessionOpen } from '@/helpers/queries/session/use-mutation-session-open';

const errorMessage = (error: unknown): string => {
  if (typeof error === 'object' && error && 'data' in error) {
    const data = (error as { data?: { message?: string } }).data;
    if (data?.message) return data.message;
  }
  return error instanceof Error ? error.message : 'Не удалось запустить лабораторную сессию';
};

export default function SessionPage() {
  const { sessionId } = useParams();
  const navigate = useNavigate();
  const ensureStarted = useRef(false);
  const [checkerPolling, setCheckerPolling] = useState(false);
  const ensure = useMutationSessionEnsure();
  const sessionQuery = useQuerySessionGet({ sessionId }, ensure.isSuccess, checkerPolling);
  const stop = useMutationSessionStop();
  const check = useMutationSessionCheck();
  const open = useMutationSessionOpen();

  useEffect(() => {
    if (!sessionId || ensureStarted.current) return;
    ensureStarted.current = true;
    ensure.mutate({ attemptId: sessionId });
  }, [ensure, sessionId]);

  useEffect(() => {
    const session = sessionQuery.data?.session;
    if (checkerPolling && session && !session.checkerRunning && session.checkerSuccessful !== undefined) {
      setCheckerPolling(false);
    }
  }, [checkerPolling, sessionQuery.data?.session]);

  if (!sessionId) {
    return <ErrorModal title='Некорректная сессия' description='В адресе отсутствует идентификатор попытки.' />;
  }

  if (ensure.isError) {
    return (
      <ErrorModal title='Сессия не запущена' description={errorMessage(ensure.error)}>
        <Button
          color='primary'
          variant='light'
          onPress={() => {
            ensureStarted.current = true;
            ensure.mutate({ attemptId: sessionId });
          }}
        >
          Повторить
        </Button>
      </ErrorModal>
    );
  }

  const session = sessionQuery.data?.session ?? ensure.data?.session;
  const phase = session?.phase ?? 'pending';
  const terminal = phase === 'ready' || phase === 'failed' || phase === 'degraded';
  const queryError = sessionQuery.isError ? errorMessage(sessionQuery.error) : undefined;

  return (
    <main className='min-h-screen flex items-center justify-center p-6 bg-background'>
      <Card className='w-full max-w-2xl'>
        <CardHeader className='flex flex-col items-start gap-2 px-6 pt-6'>
          <div className='flex w-full items-center justify-between gap-4'>
            <h1 className='text-xl font-semibold'>{session?.title || 'Лабораторная сессия'}</h1>
            <Chip color={phase === 'ready' ? 'success' : phase === 'failed' ? 'danger' : 'primary'} variant='flat'>
              {phase}
            </Chip>
          </div>
          <p className='text-small text-default-500'>Сессия: {sessionId}</p>
        </CardHeader>
        <CardBody className='gap-5 px-6'>
          {!terminal && <Progress isIndeterminate size='sm' aria-label='Запуск лаборатории' />}
          <div className='grid gap-3 sm:grid-cols-2'>
            <div className='rounded-medium border border-divider p-4'>
              <div className='font-medium'>Топология</div>
              <div className='text-small text-default-500'>{session?.topologyReady ? 'Готова' : 'Разворачивается'}</div>
            </div>
            <div className='rounded-medium border border-divider p-4'>
              <div className='font-medium'>JupyterLab</div>
              <div className='text-small text-default-500'>{session?.workspaceReady ? 'Готов' : 'Запускается'}</div>
            </div>
          </div>
          {(session?.message || queryError) && <p className='text-danger'>{session?.message || queryError}</p>}
          {session?.checkerRunning && <p className='text-small text-default-500'>Проверка лаборатории выполняется…</p>}
          {session?.checkerSuccessful !== undefined && !session.checkerRunning && (
            <p className={session.checkerSuccessful ? 'text-success' : 'text-danger'}>
              {session.checkerSuccessful
                ? 'Последняя проверка завершилась успешно.'
                : 'Последняя проверка завершилась с ошибкой.'}
            </p>
          )}
          {(check.isError || open.isError) && <p className='text-danger'>{errorMessage(check.error || open.error)}</p>}
        </CardBody>
        <CardFooter className='flex flex-wrap justify-end gap-3 px-6 pb-6'>
          <Button
            variant='flat'
            startContent={<RefreshCw className='h-4 w-4' />}
            onPress={() => sessionQuery.refetch()}
            isLoading={sessionQuery.isFetching}
          >
            Обновить
          </Button>
          <Button
            color='danger'
            variant='flat'
            startContent={<Square className='h-4 w-4' />}
            isLoading={stop.isPending}
            onPress={() => stop.mutate({ sessionId }, { onSuccess: () => navigate(RoutesLocation.home()) })}
          >
            Остановить
          </Button>
          {session?.topologyReady && (
            <Button
              as={Link}
              to={RoutesLocation.sessionTopology(sessionId)}
              variant='flat'
              startContent={<Network className='h-4 w-4' />}
            >
              Топология
            </Button>
          )}
          {phase === 'ready' && (
            <Button
              variant='flat'
              color='secondary'
              startContent={<CheckCircle className='h-4 w-4' />}
              isLoading={check.isPending || session?.checkerRunning}
              onPress={() =>
                check.mutate(
                  { sessionId },
                  {
                    onSuccess: () => {
                      setCheckerPolling(true);
                      sessionQuery.refetch();
                    }
                  }
                )
              }
            >
              Проверить
            </Button>
          )}
          {session?.workspaceUrl && (
            <Button
              color='primary'
              startContent={<ExternalLink className='h-4 w-4' />}
              isLoading={open.isPending}
              onPress={() =>
                open.mutate(
                  { sessionId },
                  {
                    onSuccess: (result) => {
                      if (result.url) window.location.assign(result.url);
                    }
                  }
                )
              }
            >
              Открыть JupyterLab
            </Button>
          )}
        </CardFooter>
      </Card>
    </main>
  );
}
