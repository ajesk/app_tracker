import { Tr, Td } from '@chakra-ui/react';

const ApplicationRow = ({ application }) => (
    <Tr key={application.id}>
        <Td>{application.title}</Td>
        <Td>{application.company}</Td>
    </Tr>
);

export default ApplicationRow;