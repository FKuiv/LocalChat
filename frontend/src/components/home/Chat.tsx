import { FC, useEffect, useState } from "react";
import { useParams } from "react-router";
import { ActionIcon, Flex, TextInput, Tooltip } from "@mantine/core";
import { Group, defaultGroup } from "@/types/group";
import { getGroupById } from "@/api/group";
import { Message, MessageRequest } from "@/types/message";
import { getMessagesByGroup } from "@/api/message";
import { IconArrowUp } from "@tabler/icons-react";
import { SendJsonMessage } from "react-use-websocket/dist/lib/types";

type chatProps = {
  sendJsonMessage: SendJsonMessage;
};

const Chat: FC<chatProps> = (props) => {
  const [group, setGroup] = useState<Group>(defaultGroup);
  const [messages, setMessages] = useState<Message[]>([]);
  const params = useParams();

  useEffect(() => {
    getGroupById(params.chatId).then((res: Group) => {
      setGroup(res);
    });

    getMessagesByGroup(params.chatId, 50).then((res: Message[]) => {
      setMessages(res);
    });
  }, [params.chatId]);

  return (
    <Flex direction="column" h="100%">
      <Flex style={{ borderBottom: "1px solid white", flexBasis: "10%" }}>
        <h3>{group.name}</h3>
      </Flex>
      <ChatMessages group={group} messages={messages} />
      <ChatInput
        sendJsonMessage={props.sendJsonMessage}
        groupId={params.chatId}
      />
    </Flex>
  );
};

type chatMessagesType = {
  group: Group;
  messages: Message[];
};

const ChatMessages: FC<chatMessagesType> = (props) => {
  return (
    <Flex direction="column" style={{ flexGrow: 1 }}>
      {props.messages.map((message) => {
        return <SingleChatMessage key={message.id} {...message} />;
      })}
    </Flex>
  );
};

const SingleChatMessage: FC<Message> = (message) => {
  return <Flex style={{ border: "1px solid blue" }}>{message.content}</Flex>;
};

const ChatInput: FC<{
  sendJsonMessage: SendJsonMessage;
  groupId: string | undefined;
}> = (props) => {
  const [newMessage, setNewMessage] = useState("");

  const handleClick = () => {
    props.sendJsonMessage<MessageRequest>({
      content: newMessage,
      group_id: props.groupId,
    });
  };

  return (
    <Flex
      direction="row"
      align="center"
      justify="space-around"
      style={{ borderTop: "1px solid white", flexBasis: "8%" }}
    >
      <TextInput
        placeholder="Send a message"
        value={newMessage}
        onChange={(event) => setNewMessage(event.currentTarget.value)}
      />
      <Tooltip label="Send">
        <ActionIcon onClick={handleClick}>
          <IconArrowUp />
        </ActionIcon>
      </Tooltip>
    </Flex>
  );
};

export default Chat;
