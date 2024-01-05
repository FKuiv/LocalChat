import { Button } from "@mantine/core";
import { logoutUser } from "@/api/user";

type logoutButtonProps = {
  fullWidth?: boolean;
};
const LogoutButton = (props: logoutButtonProps) => {
  const handleLogout = () => {
    logoutUser().then(() => {
      window.location.reload();
    });
  };

  return (
    <Button
      onClick={handleLogout}
      variant="outline"
      color="violet"
      fullWidth={props.fullWidth}
    >
      Logout
    </Button>
  );
};
export default LogoutButton;
