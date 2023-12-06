import { FC } from "react";
import { Carousel } from "@mantine/carousel";

const UsersCarousel: FC = () => {
  return (
    <Carousel
      className="carousel"
      slideSize="70%"
      height={200}
      slideGap="md"
      controlSize={40}
      loop
    >
      <div>THIGON</div>
      <div>THIGON</div>
    </Carousel>
  );
};

export default UsersCarousel;
