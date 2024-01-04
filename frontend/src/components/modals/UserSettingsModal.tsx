import {
  Modal,
  Stack,
  FileInput,
  Button,
  Title,
  TextInput,
  PasswordInput,
} from "@mantine/core";
import { RootState } from "@/redux/store";
import { useDispatch, useSelector } from "react-redux";
import { setModalOpen } from "@/redux/userSlice";
import Cookie from "universal-cookie";
import { useEffect, useState } from "react";
import UserAvatar from "../ui/UserAvatar";
import { getUserById } from "@/api/user";
import { User } from "@/types/user";
import { useForm } from "@mantine/form";

type formValues = {
  picture: File | null;
  username: string;
  password: string;
};

const UserSettingsModal = () => {
  const modalOpen = useSelector((state: RootState) => state.user.isModalOpen);
  const dispatch = useDispatch();
  const cookies = new Cookie();
  const userId = cookies.get("UserId") as string;
  const [user, setUser] = useState<User>();
  const form = useForm<formValues>({
    initialValues: {
      username: "",
      password: "",
      picture: null,
    },

    //TODO only check if the fields have actually been used
    validate: {
      username: (value) => {
        if (!form.isTouched("username")) return null;
        if (value.length < 2 || value.length > 20) {
          return "Username must be between 2-20 characters";
        }
        if (value === "Me") {
          return "Username cannot be Me";
        }
        return null;
      },
      password: (value) => {
        if (!form.isTouched("password")) return null;
        if (value.length < 4 || value.length > 50) {
          return "Password must be between 2-20 characters";
        }
        return null;
      },
      picture: (value) => {
        if (!form.isTouched("picture")) return null;
        if (value) {
          return value.size > 20000000 ? "File is too big" : null;
        }
        return "File is required";
      },
    },
  });

  useEffect(() => {
    if (!user) {
      getUserById(userId).then((res: User) => {
        setUser(res);
      });
    }
  }, [userId, user]);

  const handleSubmit = (values: formValues) => {
    console.log("new values", values);
  };

  return (
    <Modal
      opened={modalOpen}
      onClose={() => dispatch(setModalOpen(false))}
      overlayProps={{
        backgroundOpacity: 0.8,
        blur: 5,
      }}
      transitionProps={{ transition: "slide-up" }}
    >
      <form onSubmit={form.onSubmit((values) => handleSubmit(values))}>
        <Stack>
          <Stack align="center">
            <UserAvatar userId={userId} altName="Me" size="xl" />
            <Title order={2}>{user?.username}</Title>
            <FileInput
              accept="image/png,image/jpeg"
              label="Change profile picture"
              placeholder="Upload new picture"
              w="80%"
              clearable
              {...form.getInputProps("picture")}
            />
            <TextInput
              label="Change username"
              placeholder="New username"
              w="80%"
              {...form.getInputProps("username")}
            />
            <PasswordInput
              label="Change password"
              placeholder="New password"
              w="80%"
              {...form.getInputProps("password")}
            />
          </Stack>
          <Button type="submit" variant="light">
            Save
          </Button>
        </Stack>
      </form>
    </Modal>
  );
};
export default UserSettingsModal;
