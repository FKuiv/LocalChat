import { FC } from "react";
import { logoutUser } from "@/api/user";
import { Button } from "@mantine/core";

const SettingsPage: FC = () => {
  const handleLogout = () => {
    logoutUser().then(() => {
      window.location.reload();
    });
  };
  return (
    <div>
      <Button onClick={handleLogout}>Logout</Button>
    </div>
  );
};

export default SettingsPage;
