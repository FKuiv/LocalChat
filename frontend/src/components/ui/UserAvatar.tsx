import { Avatar } from "@mantine/core";
import { useEffect, useState } from "react";
import { getUserPicture } from "@/api/user";

type userAvatarProps = {
  userId: string;
  altName: string;
  picUrl?: string;
};
const UserAvatar = (props: userAvatarProps) => {
  const [picUrl, setPicUrl] = useState<string>();

  useEffect(() => {
    if (props.picUrl) {
      setPicUrl(props.picUrl);
    } else if (props.userId != "") {
      getUserPicture(props.userId).then((res: string) => {
        setPicUrl(res);
      });
    }
  }, [props]);

  return <Avatar src={picUrl} alt={props.altName} radius="md" size="lg" />;
};

export default UserAvatar;
