import { Flex } from "@chakra-ui/react"
import { Route, Routes } from "react-router-dom";
import Applications from './Applications/Applications'
import Companies from './Companies/Companies'

const Content = () => {

    return (
        <Flex margin="2em" direction="column">
            <Routes>
                <Route path="/applications" element={<Applications />} />
                <Route path="/companies" element={<Companies />} />
            </Routes>
        </Flex>
    );
}

export default Content;