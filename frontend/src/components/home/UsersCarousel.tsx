import { FC } from "react";
import { Carousel } from "@mantine/carousel";

const UsersCarousel: FC = () => {
  return (
    <div>
      <Carousel
        withIndicators
        height="100%"
        slideSize={{ base: "100%", sm: "50%", md: "33.333333%" }}
        slideGap={{ base: 0, sm: "md" }}
        loop
        align="center"
      >
        <Carousel.Slide>1</Carousel.Slide>
        <Carousel.Slide>2</Carousel.Slide>
        <Carousel.Slide>3</Carousel.Slide>
      </Carousel>
    </div>
  );
};

export default UsersCarousel;
