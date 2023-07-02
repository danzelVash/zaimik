import { footerLinks } from '@/constants';
import Link from 'next/link';
import FooterTitle from './ui/FooterTitle';

const FooterLinks: React.FC = () => {
	return (
		<div>
			<FooterTitle>Полезная информация</FooterTitle>
			<ul className='xl:mt-4 md:mt-3 mt-1.5'>
				{footerLinks?.map(obj => (
					<li key={obj.id} className='xl:mb-4 md:mb-3 mb-1.5 last:mb-0'>
						<Link
							className='hover:underline xl:text-[20px] md:text-[24px] text-[16px] leading-tight text-[#000000b3]'
							href={obj.path}
						>
							{obj.name}
						</Link>
					</li>
				))}
			</ul>
		</div>
	);
};

export default FooterLinks;
