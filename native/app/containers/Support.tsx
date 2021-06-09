import React from 'react';
 import { useSelector } from 'react-redux';
 import Support from '~/screens/Support';
 import { Auth } from '~/store/models';
 import { authSelector } from '~/store/selectors';

 export default function ConnectedSupport(): JSX.Element {
   const auth: Auth.Model = useSelector(authSelector);
   return <Support auth={auth} />;
 }
 