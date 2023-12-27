import { useContext, useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router";
import {
  ActionIcon,
  Container,
  Flex,
  TextInput,
  Title,
  Tooltip,
} from "@mantine/core";
import { Group, defaultGroup } from "@/types/group";
import { getGroupById } from "@/api/group";
import { Message } from "@/types/message";
import { getMessagesByGroup } from "@/api/message";
import { IconArrowLeft, IconArrowUp } from "@tabler/icons-react";
import { SendJsonMessage } from "react-use-websocket/dist/lib/types";
import { ReadyState } from "react-use-websocket";
import { WebSocketContext } from "@/WebSocketContext";
import Cookie from "universal-cookie";
import { nanoid } from "@reduxjs/toolkit";

const Chat = () => {
  const [group, setGroup] = useState<Group>(defaultGroup);
  const [messages, setMessages] = useState<Message[]>([]);
  const params = useParams();
  const navigate = useNavigate();
  const websocket = useContext(WebSocketContext);

  useEffect(() => {
    getGroupById(params.groupId).then((res: Group) => {
      setGroup(res);
    });

    getMessagesByGroup(params.groupId, 50).then((res: Message[]) => {
      setMessages(res);
    });
  }, [params.groupId]);

  return (
    <Flex direction="column" h="100%">
      <Flex
        align="center"
        justify="space-evenly"
        style={{ borderBottom: "1px solid white", flexBasis: "10%" }}
      >
        <ActionIcon onClick={() => navigate(-1)}>
          <IconArrowLeft />
        </ActionIcon>
        <Title m="auto">{group.name}</Title>
      </Flex>
      <ChatMessages messages={messages} />
      <ChatInput
        sendJsonMessage={websocket.sendJsonMessage}
        readyState={websocket.readyState}
        groupId={params.groupId}
      />
    </Flex>
  );
};

const ChatMessages = ({ messages }: { messages: Message[] }) => {
  return (
    <Flex direction="column" style={{ flexGrow: 1 }}>
      {messages.map((message) => {
        return <SingleChatMessage key={message.id} {...message} />;
      })}
    </Flex>
  );
};

const SingleChatMessage = (message: Message) => {
  return (
    <Container m={0} w="70%" style={{ border: "1px dashed red" }}>
      <h4 style={{ margin: 0 }}>{message.user_id}</h4>
      <p>{message.content}</p>
    </Container>
  );
};

type chatInputProps = {
  sendJsonMessage: SendJsonMessage;
  groupId: string | undefined;
  readyState: ReadyState;
};

const ChatInput = (props: chatInputProps) => {
  const [newMessage, setNewMessage] = useState("");
  const cookies = new Cookie();

  const handleClick = () => {
    const wsMessage: Message = {
      id: nanoid(),
      user_id: cookies.get("user_id"),
      group_id: props.groupId,
      content: newMessage.trim(),
      created_at: new Date(),
      updated_at: new Date(),
    };
    console.log("SENDIGN message:", wsMessage);
    props.sendJsonMessage<Message>(wsMessage);
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
        <ActionIcon
          onClick={handleClick}
          disabled={props.readyState !== ReadyState.OPEN}
        >
          <IconArrowUp />
        </ActionIcon>
      </Tooltip>
    </Flex>
  );
};

export default Chat;
