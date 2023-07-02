import FooterLinks from './FooterLinks';
import FooterLogo from './FooterLogo';
import FooterMain from './FooterMain';

const Footer: React.FC = () => {
	return (
		<footer className='md:py-10 py-5 shadow-[0px_-5px_10px_rgba(0,0,0,.25)] bg-primary-light'>
			<div className='container-xl'>
				<div className='flex xl:flex-row flex-col justify-between xl:items-center items-start gap-x-6 md:gap-y-8 gap-y-5'>
					<FooterLogo />
					<FooterMain />
					<FooterLinks />
				</div>
				<div className='xl:mt-8 md:mt-7 mt-5 text-[#000000b3] xl:text-[20px] md:text-[24px] text-[16px] leading-tight'>
					© 2023, ООО «Займик.ру». При использовании материалов гиперссылка
					на zaimik.ru обязательна. <br className='xl:block hidden' /> ИНН 7676676767, ОГРН 123456789123. 115666,
					г. Москва, проспект Вернадского, дом 666, 2 этаж.
				</div>
			</div>
		</footer>
	);
};

export default Footer;
