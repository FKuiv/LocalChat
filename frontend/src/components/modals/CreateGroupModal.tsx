import { createGroup, uploadGroupPicture } from "@/api/group";
import { getAllUsersMap } from "@/api/user";
import { Group, GroupRequest } from "@/types/group";
import {
  Button,
  Modal,
  TextInput,
  FileInput,
  Flex,
  MultiSelect,
} from "@mantine/core";
import { useForm } from "@mantine/form";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router";
import Cookie from "universal-cookie";

type createGroupModalProps = {
  opened: boolean;
  onClose: () => void;
};
type formValues = GroupRequest & {
  picture: File | null;
};

const CreateGroupModal = (props: createGroupModalProps) => {
  const [usersMap, setUsersMap] = useState<Record<string, string>>({});
  const cookies = new Cookie();
  const userId = cookies.get("UserId") as string;
  const navigate = useNavigate();
  const form = useForm<formValues>({
    initialValues: {
      name: "",
      picture: null,
      user_ids: [],
      admins: [],
      is_dm: false,
    },

    validate: {
      name: (value) => {
        return value.length === 0 ? "Name is required" : null;
      },
      picture: (value) => {
        if (value) {
          return value.size > 20000000 ? "File is too big" : null;
        }
        return "File is required";
      },
      user_ids: (value) => {
        return value.length <= 2 ? "At least 3 users are required" : null;
      },
      admins: (value) => {
        return value.length === 0 ? "At least 1 admin is required" : null;
      },
    },
  });

  useEffect(() => {
    if (Object.keys(usersMap).length === 0) {
      getAllUsersMap().then((res: Record<string, string>) => {
        setUsersMap(res);
      });
    }
    if (Object.keys(usersMap).length > 0) {
      if (form.values.user_ids.length === 0) {
        form.insertListItem(
          "user_ids",
          Object.keys(usersMap).find(
            (username) => usersMap[username] === userId
          ) as string,
          0
        );
      }
      if (form.values.admins.length === 0) {
        form.insertListItem(
          "admins",
          Object.keys(usersMap).find(
            (username) => usersMap[username] === userId
          ) as string,
          0
        );
      }
    }
  }, [usersMap, form, userId]);

  const handleFormSubmit = (values: formValues) => {
    const userIds = values.user_ids.map((user) => {
      return usersMap[user];
    });
    const adminIds = values.admins.map((admin) => {
      return usersMap[admin];
    });
    // This way the user can see the IDs because the form values are changed
    values.user_ids = userIds;
    values.admins = adminIds;
    const formData = new FormData();
    formData.append("picture", values.picture as Blob);

    createGroup({
      name: values.name,
      user_ids: values.user_ids,
      admins: values.admins,
      is_dm: values.is_dm,
    }).then((newGroup: Group) => {
      uploadGroupPicture(newGroup.id, formData).then((res) => {
        if (res.status === 200) {
          navigate(`/chat/${newGroup.id}`);
        }
      });
    });
  };

  return (
    <Modal
      opened={props.opened}
      onClose={props.onClose}
      title="Create group chat"
      transitionProps={{ transition: "slide-up" }}
    >
      <form onSubmit={form.onSubmit((values) => handleFormSubmit(values))}>
        <Flex direction="column" gap={20}>
          <TextInput
            label="Group name"
            placeholder="Group name"
            required
            {...form.getInputProps("name")}
          />
          <FileInput
            accept="image/png,image/jpeg"
            label="Group picture"
            required
            placeholder="Upload group picture"
            clearable
            {...form.getInputProps("picture")}
          />
          <MultiSelect
            label="Group members"
            placeholder="Add a user"
            data={Object.keys(usersMap)}
            searchable
            clearable
            {...form.getInputProps("user_ids")}
          />

          <MultiSelect
            label="Group admins"
            placeholder="Add a user"
            data={Object.keys(usersMap)}
            searchable
            clearable
            {...form.getInputProps("admins")}
          />

          <Button type="submit" variant="light">
            Create group
          </Button>
        </Flex>
      </form>
    </Modal>
  );
};
export default CreateGroupModal;
