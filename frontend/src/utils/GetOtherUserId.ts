import { Usernames } from "@/types/group";

export default function GetOtherUserId(
  usernameMap: Usernames,
  userId: string
): string {
  const otherUserId = Object.keys(usernameMap).filter(
    (key) => key !== userId
  )[0];
  return otherUserId;
}
