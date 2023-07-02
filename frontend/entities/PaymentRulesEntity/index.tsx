import PaymentRulesCard from '@/components/PaymentRulesCard';
import Raccoon from '@/shared/Raccoon';
import Title from '@/shared/Title';

const PaymentRulesEntity: React.FC = () => {
	return (
		<section className='section'>
			<div className='container-md'>
				<Title>
					<div className='flex items-center justify-center gap-x-1 gap-y-2 min-[360px]:flex-nowrap flex-wrap'>
						<Raccoon
							src='/static/raccoons/money_raccoon.png'
							width={91}
							height={112}
							className='shrink-0 md:hidden block'
						/>
						<div>
							<div>
								Стоит ли <span className='text-tertiary'>волноваться</span>?
							</div>
						</div>
					</div>
				</Title>
				<div className='flex flex-row lg:gap-9 gap-4 items-center lg:mt-6 md:mt-5 mt-4'>
					<PaymentRulesCard />
					<Raccoon
						src='/static/raccoons/money_raccoon.png'
						width={243}
						height={300}
						className='shrink-0 lg:w-[243px] w-[230px] lg:h-[300px] h-[285px] md:block hidden'
					/>
				</div>
			</div>
		</section>
	);
};
export default PaymentRulesEntity;
