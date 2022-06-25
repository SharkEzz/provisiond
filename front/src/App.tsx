import { Box, Flex, Grid } from '@chakra-ui/react';
import { Route, Routes } from 'react-router-dom';
import HomePage from './pages/home';
import CreatePage from './pages/create';
import Sidebar from './components/Sidebar';
import Navbar from './components/Navbar';

function App() {
  return (
    <Flex height="100vh" flexDirection="column">
      <Navbar />
      <Grid gridTemplateColumns="250px 1fr" height="100%">
        <Sidebar />
        <Box as="main" p={4}>
          <Routes>
            <Route index element={<HomePage />} />
            <Route path="/create" element={<CreatePage />} />
          </Routes>
        </Box>
      </Grid>
    </Flex>
  );
}

export default App;
