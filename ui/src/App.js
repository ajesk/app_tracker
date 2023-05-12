import NavBar from './NavBar/NavBar';
import Content from './Content/Content';
import { Flex } from '@chakra-ui/react';
import { BrowserRouter as Router } from 'react-router-dom';

function App() {
  return (
    <Flex h="100%" direction={"row"} align="center" justify="flex-start">
        <NavBar />
        <Content />
    </Flex>
  );
}

export default App;
