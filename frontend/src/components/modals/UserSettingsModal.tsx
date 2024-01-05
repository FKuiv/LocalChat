import {
  Modal,
  Stack,
  FileInput,
  Button,
  Title,
  TextInput,
  PasswordInput,
  Accordion,
  Text,
} from "@mantine/core";
import { RootState } from "@/redux/store";
import { useDispatch, useSelector } from "react-redux";
import { setModalOpen } from "@/redux/userSlice";
import Cookie from "universal-cookie";
import { useEffect, useState } from "react";
import UserAvatar from "../ui/UserAvatar";
import {
  deleteUser,
  getUserById,
  updateUser,
  uploadUserPicture,
} from "@/api/user";
import { User } from "@/types/user";
import { useForm } from "@mantine/form";
import LogoutButton from "../login/LogoutButton";

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

    validate: {
      username: (value) => {
        if (!form.isTouched("username")) return null;
        if (value === "") return null;
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
        if (value === "") return null;
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
    if (values.username !== "" || values.password !== "") {
      updateUser({ username: values.username, password: values.password }).then(
        (res) => {
          console.log(res);
        }
      );
      form.reset();
    }
    if (values.picture) {
      const formData = new FormData();
      formData.append("picture", values.picture);
      uploadUserPicture(formData).then((res) => {
        console.log(res);
      });
      form.reset();
    }
  };

  const handleDelete = () => {
    const confirmation = window.confirm(
      "Are you sure you want to delete your account?"
    );
    if (confirmation) {
      deleteUser().then((res) => {
        if (res.status === 200) {
          window.location.reload();
        }
      });
    }
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
      <Stack>
        <Stack align="center">
          <UserAvatar userId={userId} altName="Me" size="xl" />
          <Title order={2}>{user?.username}</Title>
          <Accordion variant="separated" w="100%" defaultValue="Update">
            <Accordion.Item key="Update" value="Update">
              <Accordion.Control icon="📝">Update profile</Accordion.Control>
              <Accordion.Panel>
                <form
                  onSubmit={form.onSubmit((values) => handleSubmit(values))}
                >
                  <Stack align="center">
                    <FileInput
                      accept="image/png,image/jpeg"
                      label="Change profile picture"
                      placeholder="Upload new picture"
                      clearable
                      w="100%"
                      {...form.getInputProps("picture")}
                    />
                    <TextInput
                      label="Change username"
                      placeholder="New username"
                      w="100%"
                      {...form.getInputProps("username")}
                    />
                    <PasswordInput
                      label="Change password"
                      placeholder="New password"
                      w="100%"
                      {...form.getInputProps("password")}
                    />
                    <Button type="submit" variant="light" fullWidth>
                      Save
                    </Button>
                  </Stack>
                </form>
              </Accordion.Panel>
            </Accordion.Item>
            <Accordion.Item key="Delete" value="Delete">
              <Accordion.Control icon="❌">Delete account</Accordion.Control>
              <Accordion.Panel>
                <Stack align="center">
                  <Text>
                    By deleting your account you will lose access to all the
                    messages and groups.{" "}
                    <b style={{ color: "red" }}>
                      This will not be recoverable!
                    </b>
                  </Text>
                  <Button
                    onClick={handleDelete}
                    variant="light"
                    color="red"
                    fullWidth
                  >
                    Delete
                  </Button>
                </Stack>
              </Accordion.Panel>
            </Accordion.Item>
          </Accordion>
          <LogoutButton fullWidth />
        </Stack>
      </Stack>
    </Modal>
  );
};
export default UserSettingsModal;
