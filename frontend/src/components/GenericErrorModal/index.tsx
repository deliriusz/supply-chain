import React from 'react'
import { Button, Modal, ModalProps } from 'semantic-ui-react'

type GenericModalProps = ModalProps & {
   isPositive: boolean
}

function GenericModal(props: GenericModalProps) {
   const [open, setOpen] = React.useState(props.open)
   React.useEffect(() => {
      setOpen(props.open)
   }, [props])

   return (
      <Modal
         onClose={() => setOpen(false)}
         onOpen={() => setOpen(true)}
         open={open}
      >
         <Modal.Header>{props.header}</Modal.Header>
         <Modal.Content>
            <Modal.Description>
               {props.description}
            </Modal.Description>
         </Modal.Content>
         <Modal.Actions>
            <Button
               content="Understood"
               labelPosition='right'
               icon='checkmark'
               onClick={() => setOpen(false)}
               negative={!props.isPositive}
               positive={props.isPositive}
            />
         </Modal.Actions>
      </Modal>
   )
}

export default GenericModal
