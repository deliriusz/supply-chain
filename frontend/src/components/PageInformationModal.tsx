import React from 'react'
import { Button, Icon, Modal } from 'semantic-ui-react'

function PageInformationModal() {
   const [open, setOpen] = React.useState(false)

   return (
      <Modal
         onClose={() => setOpen(false)}
         onOpen={() => setOpen(true)}
         open={open}
         trigger={
            <a className="item" href="#">
               <Icon name="question" />
            </a>
         }
      >
         <Modal.Header>Inportant information on using this page</Modal.Header>
         <Modal.Content>
            <Modal.Description>
               <ul>
                  <li>this works only on Ropsten network</li>
               </ul>
            </Modal.Description>
         </Modal.Content>
         <Modal.Actions>
            <Button
               content="Understood"
               labelPosition='right'
               icon='checkmark'
               onClick={() => setOpen(false)}
               positive
            />
         </Modal.Actions>
      </Modal>
   )
}

export default PageInformationModal
