import { paymentMethods } from '@/constants';
import Image from 'next/image';

const PaymentRulesCard: React.FC = () => {
	return (
		<div className='main-gradient py-5 xl:px-10 px-5 rounded-3xl text-white shadow-[0px_5px_10px_rgba(0,0,0,.25)]'>
			<div className='xl:text-[48px] lg:text-[35px] md:text-[28px] text-[21px] leading-tight font-bold text-center'>
				Правила оплаты и безопасность платежей
			</div>
			<p className='mt-3 xl:text-[27px] lg:text-[25px] md:text-[20px] text-[17px] font-medium leading-snug'>
				К оплате принимаются карты МИР, VISA, MasterCard. Ваши данные надежно
				защищены. При передаче данных используется защищенное соединение.
				Оформление платежей осуществляется с использованием технологий 3D-Secure
				и в соотвествии со стандратами PCI DSS.
			</p>
			<div className='mt-5 flex flex-row items-center justify-center md:gap-6 gap-4'>
				{paymentMethods?.map(obj => {
					const [width, height] = obj.size;
					return (
						<div key={obj.id}>
							<Image src={obj.src} width={width} height={height} alt='' />
						</div>
					);
				})}
			</div>
		</div>
	);
};

export default PaymentRulesCard;
