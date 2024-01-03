import { getAllUsersMap } from "@/api/user";
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

type createGroupModalProps = {
  opened: boolean;
  onClose: () => void;
};
type formValues = {
  name: string;
  picture: File | null;
  members: string[];
  admins: string[];
};

const CreateGroupModal = (props: createGroupModalProps) => {
  const [usersMap, setUsersMap] = useState<Record<string, string>>({});
  const form = useForm<formValues>({
    initialValues: {
      name: "",
      picture: null,
      members: [],
      admins: [],
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
      members: (value) => {
        return value.length === 0 ? "Members are required" : null;
      },
      admins: (value) => {
        return value.length === 0 ? "Admins are required" : null;
      },
    },
  });

  useEffect(() => {
    if (Object.keys(usersMap).length === 0) {
      getAllUsersMap().then((res: Record<string, string>) => {
        setUsersMap(res);
      });
    }
    console.log("formvalues", form.values);
  }, [usersMap, form]);

  const handleFormSubmit = (values: formValues) => {
    console.log("values", values);
  };

  return (
    <Modal
      opened={props.opened}
      onClose={props.onClose}
      title="Create group chat"
      centered
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
            {...form.getInputProps("members")}
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
            Create
          </Button>
        </Flex>
      </form>
    </Modal>
  );
};
export default CreateGroupModal;
