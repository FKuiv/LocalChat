import { Usernames } from "@/types/group";

export default function GetOtherUsername(
  usernameMap: Usernames,
  userId: string
): string {
  const otherUser = Object.keys(usernameMap).filter((key) => key !== userId)[0];
  return usernameMap[otherUser];
}

export function GetUsernameInitials(username: string): string {
  const initials = username
    .split(" ")
    .map((name) => name[0])
    .join("");
  return initials;
}
