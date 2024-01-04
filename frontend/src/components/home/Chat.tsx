import { useContext, useEffect, useState, useRef } from "react";
import { useParams, useNavigate } from "react-router";
import {
  ActionIcon,
  Container,
  Flex,
  TextInput,
  Title,
  Tooltip,
  Stack,
} from "@mantine/core";
import { Group, Usernames, defaultGroup } from "@/types/group";
import { getGroupById, getGroupPicture } from "@/api/group";
import { Message } from "@/types/message";
import { getMessagesByGroup } from "@/api/message";
import { IconArrowLeft, IconArrowUp, IconSettings } from "@tabler/icons-react";
import { SendJsonMessage } from "react-use-websocket/dist/lib/types";
import { ReadyState } from "react-use-websocket";
import { WebSocketContext } from "@/WebSocketContext";
import Cookie from "universal-cookie";
import { nanoid } from "@reduxjs/toolkit";
import { getUserPicture } from "@/api/user";
import GetOtherUsername from "@/utils/GetOtherUsername";
import GetOtherUserId from "@/utils/GetOtherUserId";
import UserAvatar from "../ui/UserAvatar";

const Chat = () => {
  const [group, setGroup] = useState<Group>(defaultGroup);
  const [messages, setMessages] = useState<Message[]>([]);
  const [fetchNewMessages, setFetchNewMessages] = useState<boolean>(true);
  const [picUrl, setPicUrl] = useState<string>();
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
      if (responseGroup.is_dm) {
        responseGroup.name = GetOtherUsername(responseGroup.usernames, userId);
        getUserPicture(GetOtherUserId(responseGroup.usernames, userId)).then(
          (res: string) => {
            setPicUrl(res);
          }
        );
      } else {
        getGroupPicture(responseGroup.id).then((res: string) => {
          setPicUrl(res);
        });
      }
      setGroup(responseGroup);
    });

    //* make sure to make this request once when first loading the page so you don't get duplicate IDs. If you make a second one than make sure to do something about websocket history
    // TODO: make sure to make this request when the user scrolls to the top of the page
    // TODO: after a certain amount of inactivity, make sure to make this request again and clear webscoket message history
    if (fetchNewMessages) {
      getMessagesByGroup(params.groupId, 50).then((res: Message[]) => {
        setMessages(res);
        setFetchNewMessages(false);
      });
    }
  }, [params.groupId, userId, messages.length, fetchNewMessages]);

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
        <Flex direction="row" align="center" justify="center" gap={10}>
          <UserAvatar
            userId=""
            altName={
              group.is_dm
                ? group.usernames[GetOtherUserId(group.usernames, userId)]
                : group.name
            }
            picUrl={picUrl}
          />
          <Title order={3}>{group.name}</Title>
        </Flex>
        <ActionIcon>
          <IconSettings />
        </ActionIcon>
      </Flex>

      <ChatMessages
        group={group}
        newMessages={fetchNewMessages}
        messages={messages.concat(websocket.messageHistory)}
        userId={userId}
        picUrl={picUrl}
      />
      <ChatInput
        sendJsonMessage={websocket.sendJsonMessage}
        readyState={websocket.readyState}
        groupId={params.groupId as string}
      />
    </Flex>
  );
};

type chatMessagesProps = {
  messages: Message[];
  group: Group;
  userId: string;
  newMessages: boolean;
  picUrl: string | undefined;
};

const ChatMessages = (props: chatMessagesProps) => {
  const messagesEndRef = useRef<null | HTMLDivElement>(null);

  useEffect(() => {
    if (messagesEndRef.current) {
      messagesEndRef.current.scrollIntoView({ behavior: "smooth" });
    }
  }, [props.newMessages]);

  return (
    <Flex direction="column" style={{ flexGrow: 1, overflow: "scroll" }}>
      <Stack
        h="50%"
        p={10}
        align="center"
        style={{ textAlign: "center", borderBottom: "1px dashed gray" }}
      >
        <UserAvatar
          userId=""
          picUrl={props.picUrl}
          altName={
            props.group.is_dm
              ? props.group.usernames[
                  GetOtherUserId(props.group.usernames, props.userId)
                ]
              : props.group.name
          }
        />

        <Title>
          {props.group.is_dm
            ? props.group.usernames[
                GetOtherUserId(props.group.usernames, props.userId)
              ]
            : props.group.name}
        </Title>
        {/* Display some stats about the group, like how many members. If it's a DM, then like last online or something */}
      </Stack>
      {props.messages.map((message) => {
        if (message.group_id === props.group.id) {
          return (
            <SingleChatMessage
              key={message.id}
              message={message}
              usernameMap={props.group.usernames}
              userId={props.userId}
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
  usernameMap,
  userId,
}: {
  message: Message;
  usernameMap: Usernames;
  userId: string;
}) => {
  const [username, setUsername] = useState<string>("Me");
  const [containerStyle, setContainerStyle] = useState({
    border: "1px dashed blue",
    alignSelf: "flex-end",
  });
  useEffect(() => {
    if (userId !== message.user_id) {
      setUsername(usernameMap[message.user_id]);
      setContainerStyle({ border: "1px dashed red", alignSelf: "flex-start" });
    }
  }, [userId, message.user_id, usernameMap]);

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
    props.sendJsonMessage<Message>(wsMessage);
  };

  return (
    <Flex
      direction="row"
      align="center"
      justify="space-around"
      style={{ borderTop: "1px solid white", flexBasis: "10%" }}
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
