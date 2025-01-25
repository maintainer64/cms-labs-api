import { InputOtp } from '@nextui-org/input-otp';
import { useState } from 'react';
import { Button, Card, CardBody, CardFooter, CardHeader } from '@nextui-org/react';
import { Divider } from '@nextui-org/divider';

export type Teammate = {
  avatar: string;
  name: string;
};

type TeammatesListProps = {
  teammates: Teammate[];
};

export const TeammatesList = ({ teammates }: TeammatesListProps) => {
  return (
    <div className="mt-4">
      {teammates.map((teammate, index) => (
        <img
          key={index}
          src={teammate.avatar}
          alt={teammate.name}
          className="w-12 h-12 rounded-full mr-2"
        />
      ))}
    </div>
  );
};

export const MyComponent = () => {
  const [roomNumber, setRoomNumber] = useState<string>('');
  const [teammates, setTeammates] = useState<Teammate[]>([]);

  const handleTeammateConnect = (teammate: Teammate) => {
    setTeammates([...teammates, teammate]);
  };

  return (
    <Card className="max-w-[800px]">
      <CardHeader className="flex gap-3">
        Waiting Room #{roomNumber}
      </CardHeader>
      <Divider />
      <CardBody>
        <div className="flex flex-col justify-center items-center">
          <InputOtp
            length={6}
            value={roomNumber}
            onValueChange={setRoomNumber}
          />
        </div>
      </CardBody>
      <Divider />
      <CardBody>
        <TeammatesList teammates={teammates} />
      </CardBody>
      <Divider />
      <CardFooter>
        <Button variant="flat" color="primary">Enter Room</Button>
      </CardFooter>
    </Card>
  );
};