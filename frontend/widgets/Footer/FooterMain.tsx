import { general } from '@/constants';
import FooterTitle from './ui/FooterTitle';

const FooterMain: React.FC = () => {
	return (
		<div>
			<FooterTitle>
				Служба поддержки всегда на <span className='text-tertiary'>связи</span>!
			</FooterTitle>
			<div className='mt-1'>
				<a
					className='font-bold xl:text-[26px] md:text-[32px] text-[20px] transition-opacity duration-200 hover:opacity-70'
					href={`mailto:${general.email}`}
				>
					{general.email}
				</a>
			</div>
			<p className='md:text-[24px] text-[16px] leading-tight mt-1'>
				Ответим на ваш вопрос в любое время!
			</p>
		</div>
	);
};

export default FooterMain;
