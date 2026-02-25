import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
import { Loading } from '@/components/scroll/loader';
import { useQueryStaticDocsGet } from '@/helpers/queries/docs/static-docs';

interface Props {
  path: string;
  variables?: Record<string, string>;
}

const MarkdownViewer = ({ path, variables = {} }: Props) => {
  // Функция замены переменных в контенте
  const queryUseDocs = useQueryStaticDocsGet({ path });
  const replaceVariables = (text: string) => {
    let result = text;
    for (const [key, value] of Object.entries(variables)) {
      result = result.replace(new RegExp(`\\{\\{${key}\\}\\}`, 'g'), value);
    }
    return result;
  };

  if (queryUseDocs.isLoading) return <Loading size='md' />;
  return (
    <div
      className='prose prose-slate dark:prose-invert
  prose-h1:font-bold prose-h1:text-xl
  prose-a:text-blue-600 prose-p:text-justify prose-img:rounded-xl prose-lg max-w-none'
    >
      <ReactMarkdown remarkPlugins={[remarkGfm]}>{replaceVariables(queryUseDocs.data ?? '')}</ReactMarkdown>
    </div>
  );
};

export default MarkdownViewer;
