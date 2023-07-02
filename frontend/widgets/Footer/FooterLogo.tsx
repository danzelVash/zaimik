import Image from 'next/image';

const FooterLogo: React.FC = () => {
	return (
		<div className='xl:w-auto md:w-[200px] w-[115px]'>
			<Image src='/logo_big.png' width={284} height={238} alt='' />
		</div>
	);
};

export default FooterLogo;
