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
import { IconArrowLeft, IconArrowUp, IconSettings } from "@tabler/icons-react";
import { SendJsonMessage } from "react-use-websocket/dist/lib/types";
import { ReadyState } from "react-use-websocket";
import { WebSocketContext } from "@/WebSocketContext";
import Cookie from "universal-cookie";
import { nanoid } from "@reduxjs/toolkit";
import { getUsername } from "@/api/user";

const Chat = () => {
  const [group, setGroup] = useState<Group>(defaultGroup);
  const [messages, setMessages] = useState<Message[]>([]);
  const params = useParams();
  const navigate = useNavigate();
  const websocket = useContext(WebSocketContext);

  useEffect(() => {
    // TODO: redirect user if this happens
    if (!params.groupId) {
      console.error("No group id provided");
      return;
    }
    getGroupById(params.groupId).then((res: Group) => {
      setGroup(res);
    });

    //* make sure to make this request once when first loading the page so you don't get duplicate IDs. If you make a second one than make sure to do something about websocket history
    // TODO: make sure to make this request when the user scrolls to the top of the page
    // TODO: after a certain amount of inactivity, make sure to make this request again and clear webscoket message history
    if (messages.length < 50) {
      getMessagesByGroup(params.groupId, 50).then((res: Message[]) => {
        console.log("Making request to get messages by group");
        setMessages(res);
      });
    }
  }, [params.groupId, messages.length]);

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
        <ActionIcon>
          <IconSettings />
        </ActionIcon>
      </Flex>
      <ChatMessages
        groupId={params.groupId as string}
        messages={messages}
        wsMessages={websocket.messageHistory}
      />
      <ChatInput
        sendJsonMessage={websocket.sendJsonMessage}
        readyState={websocket.readyState}
        groupId={params.groupId as string}
      />
    </Flex>
  );
};

const ChatMessages = ({
  messages,
  wsMessages,
  groupId,
}: {
  messages: Message[];
  wsMessages: Message[];
  groupId: string;
}) => {
  return (
    <Flex direction="column" style={{ flexGrow: 1, overflow: "scroll" }}>
      {messages.map((message) => {
        if (message.group_id === groupId) {
          return <SingleChatMessage key={message.id} {...message} />;
        }
      })}

      {wsMessages.map((message) => {
        if (message.group_id === groupId) {
          return <SingleChatMessage key={message.id} {...message} />;
        }
      })}
    </Flex>
  );
};

const SingleChatMessage = (message: Message) => {
  const cookies = new Cookie();
  const [username, setUsername] = useState<string>("");
  let containerStyle = { border: "1px dashed red", alignSelf: "flex-start" };

  if (cookies.get("UserId") !== message.user_id) {
    containerStyle = { border: "1px dashed blue", alignSelf: "flex-end" };
  }
  getUsername(message.user_id).then((res) => {
    console.log("username:", res);
    setUsername(res);
  });

  return (
    <Container m={0} w="70%" style={containerStyle}>
      <h4 style={{ margin: 0 }}>{username}</h4>
      <p>{message.content}</p>
    </Container>
  );
};

type chatInputProps = {
  sendJsonMessage: SendJsonMessage;
  groupId: string;
  readyState: ReadyState;
};

const ChatInput = (props: chatInputProps) => {
  const [newMessage, setNewMessage] = useState("");
  const cookies = new Cookie();

  const handleClick = () => {
    const wsMessage: Message = {
      id: nanoid(),
      user_id: cookies.get("UserId"),
      group_id: props.groupId,
      content: newMessage.trim(),
      created_at: new Date(),
      updated_at: new Date(),
      DeletedAt: null,
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
