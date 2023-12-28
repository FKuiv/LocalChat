import { useContext, useEffect, useState, useRef } from "react";
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
  const cookies = new Cookie();
  const userId = cookies.get("UserId");

  useEffect(() => {
    if (!params.groupId) {
      console.error("No group id provided");
      return;
    }
    getGroupById(params.groupId).then((responseGroup: Group) => {
      if (responseGroup.isDm) {
        const otherUserId = responseGroup.users.filter(
          (user) => user.id !== userId
        )[0].id;
        getUsername(otherUserId).then((username: string) => {
          responseGroup.name = username;
          setGroup(responseGroup);
        });
      } else {
        setGroup(responseGroup);
      }
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
  }, [params.groupId, messages.length, userId]);

  return (
    <Flex direction="column" h="100%">
      <Flex
        align="center"
        justify="space-around"
        style={{ borderBottom: "1px solid white", flexBasis: "10%" }}
      >
        <ActionIcon onClick={() => navigate(-1)}>
          <IconArrowLeft />
        </ActionIcon>
        <Title>{group.name}</Title>
        <ActionIcon>
          <IconSettings />
        </ActionIcon>
      </Flex>
      <ChatMessages
        groupId={params.groupId as string}
        messages={messages}
        wsMessages={websocket.messageHistory}
        userId={userId}
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
  userId,
}: {
  messages: Message[];
  wsMessages: Message[];
  groupId: string;
  userId: string;
}) => {
  const messagesEndRef = useRef<null | HTMLDivElement>(null);

  useEffect(() => {
    if (messagesEndRef.current) {
      messagesEndRef.current.scrollIntoView({ behavior: "smooth" });
    }
  }, [messages]);

  return (
    <Flex direction="column" style={{ flexGrow: 1, overflow: "auto" }}>
      {messages.map((message) => {
        if (message.group_id === groupId) {
          return (
            <SingleChatMessage
              key={message.id}
              message={message}
              userId={userId}
            />
          );
        }
      })}

      {wsMessages.map((message) => {
        if (message.group_id === groupId) {
          return (
            <SingleChatMessage
              key={message.id}
              message={message}
              userId={userId}
            />
          );
        }
      })}
      <div ref={messagesEndRef} />
    </Flex>
  );
};

const SingleChatMessage = ({
  message,
  userId,
}: {
  message: Message;
  userId: string;
}) => {
  const [username, setUsername] = useState<string>("Me");
  let containerStyle = { border: "1px dashed blue", alignSelf: "flex-end" };

  if (userId !== message.user_id) {
    getUsername(message.user_id).then((res) => {
      setUsername(res);
    });
    containerStyle = { border: "1px dashed red", alignSelf: "flex-start" };
  }

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
