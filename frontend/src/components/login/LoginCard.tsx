import { TextInput, Button, Group, Box } from "@mantine/core";
import { hasLength, useForm } from "@mantine/form";

export default function LoginCard() {
  const form = useForm({
    initialValues: {
      username: "",
      password: "",
    },

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
    <Box maw={340} mx="auto">
      <form onSubmit={form.onSubmit((values) => console.log(values))}>
        <TextInput
          withAsterisk
          label="Username"
          {...form.getInputProps("username")}
        />

        <TextInput
          withAsterisk
          label="Password"
          {...form.getInputProps("password")}
        />

        <Group justify="flex-end" mt="md">
          <Button type="submit">Login</Button>
        </Group>
      </form>
    </Box>
  );
}
