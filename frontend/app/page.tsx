import ToTop from '@/widgets/ToTop';
import IntroEntity from '@/entities/IntroEntity';
import LoanConditionEntity from '@/entities/LoanConditionEntity';
import ObtainingStagesEntity from '@/entities/ObtainingStagesEntity';
import PaymentRulesEntity from '@/entities/PaymentRulesEntity';

export const metadata = {
	title: 'Главная'
};

const HomePage = () => {
	return <>
		<ToTop />
		<IntroEntity />
		<LoanConditionEntity />
		<ObtainingStagesEntity />
		<PaymentRulesEntity />
	</>;
};

export default HomePage;
