import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Login from './components/Login';
import Payment from './components/Payment';
<<<<<<< HEAD
import Equipment from './components/Equipment';
=======
import Promotion from './components/Promotion';
>>>>>>> 0d53127f (ทำ frontend ของระบบโปรโมชั่น - close #76)

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/Payment', Payment);
<<<<<<< HEAD
    router.registerRoute('/Equipment', Equipment);
=======
    router.registerRoute('/Promotion', Promotion);
>>>>>>> 0d53127f (ทำ frontend ของระบบโปรโมชั่น - close #76)
  },
});