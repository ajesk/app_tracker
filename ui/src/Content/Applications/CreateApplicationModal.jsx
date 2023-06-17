import {
    Modal, ModalBody, ModalContent,
    ModalOverlay, ModalHeader, ModalCloseButton,
} from "@chakra-ui/react";
import { createApplication } from "../../api/applictions";
import CreateApplicationForm from "./CreateApplicationForm";

const CreateApplicationModal = ({ isOpen, handleClose }) => {
    const handleSubmit = async (application) => {
        if (application.max_salary) application.max_salary = parseInt(application.max_salary);
        if (application.min_salary) application.min_salary = parseInt(application.min_salary);
        const companyId = application.company;
        delete application.company;
        
        console.log(application)
        await createApplication(1, companyId, application);
        handleClose();
    };

    return (
        <Modal isOpen={isOpen} onClose={handleClose}>
            <ModalOverlay />
            <ModalContent>
                <ModalHeader>Create Application</ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                    <CreateApplicationForm handleSubmit={handleSubmit} />
                </ModalBody>
            </ModalContent>
        </Modal>
    )
};

export default CreateApplicationModal;