import { FC } from "react";
import { Carousel } from "@mantine/carousel";

const UsersCarousel: FC = () => {
  return (
    <Carousel
      className="carousel border"
      slideSize="70%"
      height={200}
      slideGap="md"
      controlSize={40}
      loop
    >
      <Carousel.Slide>Thing</Carousel.Slide>
    </Carousel>
  );
};

export default UsersCarousel;
