import { FC, useState } from "react";
import { AppShell, Burger, Button, FileInput, Image } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import Navbar from "@/components/navigation/Navbar";
import Logo from "@/components/ui/Logo";
import { UserEndpoints, api } from "@/endpoints";
import UsersCarousel from "@/components/home/UsersCarousel";

const HomePage: FC = () => {
  const [opened, { toggle }] = useDisclosure();
  const [profilePic, setProfilePic] = useState<File | null>(null);
  const [picUrl, setpicUrl] = useState("");

  const handleLogout = () => {
    api
      .get(UserEndpoints.logout)
      .then((res) => {
        console.log("logout res", res);
        window.location.reload();
      })
      .catch((err) => {
        console.log("logout err", err);
      });
  };

  const handleUpload = () => {
    const formData = new FormData();
    formData.append("picture", profilePic as Blob);
    api
      .post(UserEndpoints.profilepic, formData, {
        headers: { "Content-Type": "multipart/form-data" },
      })
      .then((res) => {
        console.log("res from file uplaod", res);
      })
      .catch((err) => {
        console.log("err from file uploa", err);
      });
  };

  const handlePic = () => {
    api
      .get(UserEndpoints.profilepic)
      .then((res) => {
        console.log("Res geting pic", res);
        setpicUrl(res.data);
      })
      .catch((err) => {
        console.log("err gettign pic", err);
      });
  };

  return (
    <AppShell
      header={{ height: 60 }}
      navbar={{ width: 300, breakpoint: "sm", collapsed: { mobile: !opened } }}
    >
      <AppShell.Header className="header">
        <Burger opened={opened} onClick={toggle} hiddenFrom="sm" size="sm" />
        <Logo />
        <UsersCarousel />
        <Button onClick={handleLogout}>Logout</Button>
      </AppShell.Header>
      <AppShell.Navbar p="md">
        <Navbar />
      </AppShell.Navbar>
      walkthrough
      <AppShell.Main>
        <FileInput
          label="Profile pic"
          description="Upload your profile picture"
          placeholder="my_profile_picture.png"
          value={profilePic}
          onChange={setProfilePic}
        />
        <Button onClick={handleUpload}>Upload</Button>
        <Button onClick={handlePic}>Get Profile pic</Button>
        <Image radius="md" src={picUrl} />
      </AppShell.Main>
    </AppShell>
  );
};

export default HomePage;
