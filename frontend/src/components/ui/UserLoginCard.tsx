import { Login, defaultLogin } from "@/types/Login";
import { TextInput, Flex, Button, Group, Box } from "@mantine/core";
import { hasLength, useForm } from "@mantine/form";
import { FC } from "react";

type userLoginCardType = {
  onSubmit: ({ username, password }: Login) => void;
  buttonLabel: string;
};

const UserLoginCard: FC<userLoginCardType> = (props) => {
  const { onSubmit, buttonLabel } = props;

  const form = useForm({
    initialValues: defaultLogin,

    validate: {
      username: hasLength(
        { min: 2, max: 20 },
        "Username must be between 2-20 characters"
      ),
      password: hasLength(
        { min: 4, max: 50 },
        "Password must be between 4-50 characters"
      ),
    },
  });

  return (
    <Box mx="auto" maw="80%">
      <form onSubmit={form.onSubmit((values) => onSubmit(values))}>
        <Flex direction="column" gap="xl">
          <TextInput
            w="100%"
            withAsterisk
            label="Username"
            {...form.getInputProps("username")}
          />

          <TextInput
            withAsterisk
            label="Password"
            type="password"
            {...form.getInputProps("password")}
          />

          <Group justify="flex-end" mt="md">
            <Button type="submit">{buttonLabel}</Button>
          </Group>
        </Flex>
      </form>
    </Box>
  );
};

export default UserLoginCard;
