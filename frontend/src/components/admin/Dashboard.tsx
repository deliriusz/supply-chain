import React, { useState } from "react";
import { Icon, Menu, MenuItemProps, Segment, Statistic } from "semantic-ui-react";
import AddProductPane from "./ProductSearchPane";
import StatisticsPane from "./StatisticsPane";

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
               name='add product'
               active={activeItem === 'add product'}
               onClick={handleItemClick}
            />
         </Menu>
         {activeItem == 'statistics' && <StatisticsPane />}
         {activeItem == 'add product' && <AddProductPane />}
      </>
   )
}

export default Dashboard;