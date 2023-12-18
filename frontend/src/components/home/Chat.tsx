import { FC } from "react";
import { useParams } from "react-router";
const Chat: FC = () => {
  const params = useParams();

  return <div>Specific chat with id: {params.chatId}</div>;
};
export default Chat;
