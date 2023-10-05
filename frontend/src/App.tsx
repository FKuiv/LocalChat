import { Button, TextInput } from "@mantine/core";
import "./App.styles.css";
import { useEffect, useState } from "react";

export default function App() {
  const [text, setText] = useState("");

  useEffect(() => {
    console.log(text);
  }, [text]);

  const handleClick = () => {
    alert(`The message: ${text}`);
  };

  return (
    <div>
      <TextInput
        mt="md"
        label="Send message"
        value={text}
        onChange={(event) => setText(event.currentTarget.value)}
      />
      <Button mt="lg" onClick={handleClick}>
        Send
      </Button>
    </div>
  );
}
