import { getAllUserGroups } from "@/api/group";
import { getUserPicture } from "@/api/user";
import { Group } from "@/types/group";
import GetOtherUserId from "@/utils/GetOtherUserId";
import { GetUsernameInitials } from "@/utils/GetOtherUsername";
import { Flex, Container, Avatar, Title } from "@mantine/core";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router";
import Cookie from "universal-cookie";

const ChatGroups = () => {
  const [groups, setGroups] = useState<Group[]>();

  useEffect(() => {
    getAllUserGroups().then((res: Group[]) => {
      setGroups(res);
    });
  }, []);

  return (
    <Flex w="100%" h="100%" direction="column">
      {groups?.map((group: Group) => (
        <ChatGroup {...group} key={group.id} />
      ))}
    </Flex>
  );
};

const ChatGroup = (group: Group) => {
  const navigate = useNavigate();
  const cookies = new Cookie();
  const otherUserId = GetOtherUserId(group.usernames, cookies.get("UserId"));
  const [picUrl, setPicUrl] = useState<string>();

  const handleClick = () => {
    navigate(`/chat/${group.id}`);
  };

  useEffect(() => {
    console.log("group", group);
    if (group.is_dm) {
      getUserPicture(otherUserId).then((res: string) => {
        setPicUrl(res);
      });
    } else {
      return;
    }
  }, [group, otherUserId]);

  return (
    <Container
      w="100%"
      h={80}
      style={{ borderBottom: "1px solid var(--_app-shell-border-color)" }}
      onClick={handleClick}
    >
      <Flex direction="row" align="center" gap="md" h="100%">
        <Avatar src={picUrl} alt={group.usernames[otherUserId]} size="lg">
          {GetUsernameInitials(group.usernames[otherUserId])}
        </Avatar>
        <Title order={3}>
          {group.is_dm ? group.usernames[otherUserId] : group.name}
        </Title>
      </Flex>
    </Container>
  );
};

export default ChatGroups;
