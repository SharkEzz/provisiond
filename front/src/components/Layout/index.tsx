import { Box, Container, Text } from '@chakra-ui/react';

type LayoutProps = {
  children: React.ReactNode;
  title: string;
};

export default function Layout({ title, children }: LayoutProps) {
  return (
    <Box>
      <Text as="h1" mb={3} fontSize="2xl">
        {title}
      </Text>
      <Container maxW="container.xl">{children}</Container>
    </Box>
  );
}
