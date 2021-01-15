import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Login from './components/Login';
import Payment from './components/Payment';
import Equipment from './components/Equipment';
import Bookcourse from './components/ฺฺBookcourse';
import Employee from './components/Employee';
import Promotion from './components/Promotion';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/Payment', Payment);
    router.registerRoute('/Equipment', Equipment);
    router.registerRoute('/Bookcourse', Bookcourse);
    router.registerRoute('/Employee', Employee);
    router.registerRoute('/Promotion', Promotion);
  },
});