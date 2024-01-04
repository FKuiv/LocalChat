import { createGroup, groupExistingByUserIdsAndAdmins } from "@/api/group";
import { getAllUsersMap } from "@/api/user";
import { Group } from "@/types/group";
import { Carousel } from "@mantine/carousel";
import { Stack, Text } from "@mantine/core";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router";
import Cookie from "universal-cookie";
import UserAvatar from "../ui/UserAvatar";
import { useDispatch } from "react-redux";
import { setModalOpen } from "@/redux/userSlice";

const UserCarousel = () => {
  const [usersMap, setUsersMap] = useState<Record<string, string>>({});
  const cookie = new Cookie();

  useEffect(() => {
    getAllUsersMap().then((res: Record<string, string>) => {
      setUsersMap(res);
    });
  }, []);

  return (
    <Carousel
      align="start"
      slideSize="30%"
      slideGap="0"
      pt={5}
      h="10%"
      loop
      withControls={false}
      style={{ borderBottom: "1px solid var(--_app-shell-border-color)" }}
    >
      {Object.keys(usersMap).map((username: string) => {
        if (usersMap[username] !== cookie.get("UserId")) {
          return (
            <UserProfileSlide
              key={usersMap[username]}
              otherUsername={username}
              otherUserId={usersMap[username]}
              userId={cookie.get("UserId") as string}
            />
          );
        }
        return (
          <UserProfileSlide
            key={usersMap[username]}
            otherUsername={"Me"}
            otherUserId={usersMap[username]}
            userId={cookie.get("UserId") as string}
          />
        );
      })}
    </Carousel>
  );
};

const UserProfileSlide = ({
  otherUsername,
  otherUserId,
  userId,
}: {
  otherUsername: string;
  otherUserId: string;
  userId: string;
}) => {
  const users = [userId, otherUserId];
  const navigate = useNavigate();
  const dispatch = useDispatch();

  const handleClick = () => {
    if (otherUserId != userId) {
      groupExistingByUserIdsAndAdmins(users, users).then((res: Group[]) => {
        if (res.length != 0) {
          navigate(`/chat/${res[0].id}`);
        } else {
          createGroup({
            name: `${otherUserId} ${userId}`,
            user_ids: users,
            admins: users,
            is_dm: true,
          }).then((res: Group) => {
            navigate(`/chat/${res.id}`);
          });
        }
      });
    } else {
      dispatch(setModalOpen(true));
    }
  };

  // TODO: add active badge to the user if they are connected via websocket
  return (
    <Carousel.Slide onClick={handleClick}>
      <Stack align="center" gap={5} h="100%" justify="center">
        <UserAvatar userId={otherUserId} altName={otherUsername} />
        <Text size="md">{otherUsername}</Text>
      </Stack>
    </Carousel.Slide>
  );
};

export default UserCarousel;
