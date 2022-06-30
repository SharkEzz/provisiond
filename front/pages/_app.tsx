import { Box, ChakraProvider, Flex, Grid } from '@chakra-ui/react';
import type { AppProps } from 'next/app';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <ChakraProvider>
      <Flex height="100vh" flexDirection="column">
        <Navbar />
        <Grid gridTemplateColumns="250px 1fr" height="100%">
          <Sidebar />
          <Box as="main" p={4}>
            <Component {...pageProps} />
          </Box>
        </Grid>
      </Flex>
    </ChakraProvider>
  );
}

export default MyApp;
