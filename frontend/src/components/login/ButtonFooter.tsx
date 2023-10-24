import { CardFooter } from "../ui/card";
import { Button } from "../ui/button";

export default function ButtonFooter({ label }: { label: string }) {
  return (
    <CardFooter className="justify-end">
      <Button type="submit">{label}</Button>
    </CardFooter>
  );
}
