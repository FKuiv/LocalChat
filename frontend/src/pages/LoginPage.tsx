import { Button } from "../components/ui/button";
import { Textarea } from "../components/ui/textarea";
import { useEffect, useState } from "react";

export default function LoginPage() {
  const [text, setText] = useState("");

  const handleClick = () => {
    alert(`The message: ${text}`);
  };

  useEffect(() => {
    console.log(text);
  }, [text]);

  return (
    <div className="flex flex-col justify-center items-center pt-10 space-y-10">
      <Textarea
        className="w-1/2"
        onChange={(event) => setText(event.target.value)}
      />
      <Button onClick={handleClick}>Send</Button>
    </div>
  );
}
