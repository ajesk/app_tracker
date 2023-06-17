
import React from 'react';
import { Table, Thead, Tbody, Tr, Th, TableContainer } from '@chakra-ui/react';
import ApplicationRow from './ApplicationRow';

const ApplicationTable = ({ applications = [], showApplicationModal }) => {
    return (
        <TableContainer>
            <Table variant="simple" colorScheme="gray">
                <Thead>
                    <Tr>
                        <Th>Title</Th>
                        <Th>Company</Th>
                        <Th>Salary Expectations</Th>
                        <Th>Notes</Th>
                        <Th />
                    </Tr>
                </Thead>
                <Tbody>
                    {
                        applications.map((application) => <ApplicationRow key={application.id} application={application} showApplicationModal={showApplicationModal} />)
                    }
                </Tbody>
            </Table>
        </TableContainer>
    );
};

export default ApplicationTable;
