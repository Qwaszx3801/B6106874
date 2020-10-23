import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import Button from '@material-ui/core/Button';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';

//const username = { givenuser: login}
const WelcomePage: FC<{}> = () => {
  const profile = { givenName: '' };
 
 return (
   <Page theme={pageTheme.service}>
     <Header title={` ${profile.givenName || 'ระบบทุนการศึกษา'}`}>
     
 
     </Header>
     
     <Content>
       <ContentHeader title="ระบบรับข้อมูลจากผู้ให้ทุนการศึกษา">    

         <Link component={RouterLink} to="/create">
           <Button variant="outlined" color="primary">
             Add Data
           </Button>
           </Link> 
&nbsp;&nbsp;&nbsp;&nbsp;
            <Link component={RouterLink} to="/">
           <Button variant="outlined" color="secondary">
             LOG OUT
           </Button>
       </Link> 
       </ContentHeader>
       
     </Content>
   </Page>
 );
};
 
export default WelcomePage;