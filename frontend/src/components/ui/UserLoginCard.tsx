import { Login, defaultLogin } from "@/types/user";
import { TextInput, Flex, Button, Group, Box } from "@mantine/core";
import { hasLength, useForm } from "@mantine/form";

type userLoginCardType = {
  onSubmit: ({ username, password }: Login) => void;
  buttonLabel: string;
};

const UserLoginCard = (props: userLoginCardType) => {
  const { onSubmit, buttonLabel } = props;

  const form = useForm({
    initialValues: defaultLogin,

    validate: {
      username: (value) => {
        if (value.length < 2 || value.length > 20) {
          return "Username must be between 2-20 characters";
        }
        if (value === "Me") {
          return "Username cannot be Me";
        }
        return null;
      },
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
