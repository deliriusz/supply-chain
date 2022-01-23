import React, { useState } from "react";
import { Icon, Menu, MenuItemProps, Segment, Statistic } from "semantic-ui-react";
import SalesChart from "./SalesChart";

//TODO: think about making menu sticky
const Dashboard = () => {

   const [activeItem, setActiveItem] = useState<string>('statistics');
   const handleItemClick = (event: React.MouseEvent<HTMLAnchorElement>, data: MenuItemProps) => { setActiveItem(data.name ? data.name : 'statistics') };

   return (
      <>
         <Menu pointing secondary>
            <Menu.Item
               name='statistics'
               active={activeItem === 'statistics'}
               onClick={handleItemClick}
            />
            <Menu.Item
               name='sales'
               active={activeItem === 'sales'}
               onClick={handleItemClick}
            />
            <Menu.Item
               name='product search'
               active={activeItem === 'product search'}
               onClick={handleItemClick}
            />
         </Menu>
         {activeItem}
         <Segment compact>
            <Statistic.Group>
               <Statistic>
                  <Statistic.Value>22</Statistic.Value>
                  <Statistic.Label>Saves</Statistic.Label>
               </Statistic>

               <Statistic>
                  <Statistic.Value text>
                     Three
                     <br />
                     Thousand
                  </Statistic.Value>
                  <Statistic.Label>Signups</Statistic.Label>
               </Statistic>

               <Statistic>
                  <Statistic.Value>
                     <Icon name='plane' />5
                  </Statistic.Value>
                  <Statistic.Label>Flights</Statistic.Label>
               </Statistic>

               <Statistic>
                  <Statistic.Value>
                     42
                  </Statistic.Value>
                  <Statistic.Label>Team Members</Statistic.Label>
               </Statistic>
            </Statistic.Group>
         </Segment>
         <SalesChart />
         <h3>contract a product button</h3>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
         <p>list of all products with filtering</p>
      </>
   )
}

export default Dashboard;