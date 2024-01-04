import { Avatar } from "@mantine/core";
import { useEffect, useState } from "react";
import { getUserPicture } from "@/api/user";

type userAvatarProps = {
  userId: string;
  altName: string;
  picUrl?: string;
  size?: string;
  onClick?: () => void;
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

  return (
    <Avatar
      src={picUrl}
      alt={props.altName}
      radius="md"
      size={props.size ? props.size : "lg"}
      onClick={props.onClick}
    />
  );
};

export default UserAvatar;
