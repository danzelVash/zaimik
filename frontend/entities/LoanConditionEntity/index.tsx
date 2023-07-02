import LoanConditionCard from '@/components/LoanConditionCard';
import { loanConditions } from '@/constants';
import PageTitle from '@/shared/Title';

const LoanConditionEntity: React.FC = () => {
	return (
		<section className='section'>
			<div className='container-md'>
				<PageTitle>
					Условия для <span className='text-tertiary'>получения</span> займа
				</PageTitle>
				<div className='grid grid-cols-6 xl:gap-8 md:gap-6 gap-5 lg:mt-10 md:mt-7 mt-5'>
					{loanConditions?.map(obj => (
						<LoanConditionCard key={obj.id} {...obj} />
					))}
				</div>
			</div>
		</section>
	);
};

export default LoanConditionEntity;
