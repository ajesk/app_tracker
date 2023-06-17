import { Tr, Td } from '@chakra-ui/react';
import { ViewIcon } from '@chakra-ui/icons';

const ApplicationRow = ({ application, showApplicationModal }) => (
    <Tr key={application.id}>
        <Td>{application.title}</Td>
        <Td>{application.company}</Td>
        <Td>${application.min_salary} - ${application.max_salary}</Td>
        <Td>{application.notes}</Td>
        <Td><ViewIcon onClick={() => showApplicationModal(application.id)}/></Td>
    </Tr>
);

export default ApplicationRow;