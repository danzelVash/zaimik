import { navLinks } from '@/constants';
import Image from 'next/image';
import Link from 'next/link';
import styles from './styles.module.scss';

interface IHeaderNavProps {
	isOpened: Boolean;
	hideNav: () => void;
}

const HeaderNav: React.FC<IHeaderNavProps> = ({ isOpened, hideNav }) => {
	if (!navLinks.length) return;

	return (
		<nav
			className={`${
				isOpened ? 'opacity-100 visible' : 'opacity-0 invisible'
			} lg:opacity-100 lg:visible transition-all duration-300 lg:block flex justify-end lg:absolute fixed lg:top-1/2 top-0 lg:right-1/2 lg:translate-x-1/2 lg:-translate-y-1/2 right-0 lg:w-auto w-full h-full overflow-y-auto lg:z-[2] z-[-1]`}
		>
			<ul
				className={`${styles['nav-list']} ${
					isOpened ? 'translate-x-0' : 'translate-x-full'
				} transition-all duration-300 lg:translate-x-0 h-full flex lg:flex-row flex-col lg:items-center items-start 2xl:gap-x-14 gap-x-8 lg:p-0 lg:pt-0 md:p-7 p-4 md:pt-[108px] pt-[92px]`}
			>
				{navLinks?.map(obj => {
					const [width, height] = obj.icon.size;
					return (
						<li
							key={obj.id}
							className='lg:w-auto lg:h-auto w-full lg:pb-0 lg:mb-0 lg:border-b-0 pb-3 mb-3 border-b-[3px] border-[#000000b3] last:border-b-0 last:pb-0 last:mb-0'
						>
							<Link
								onClick={hideNav}
								href={obj.path}
								className='flex flex-row items-center 2xl:gap-x-4 lg:gap-x-3 md:gap-x-5 gap-x-1 transition-opacity duration-300 hover:opacity-70'
							>
								<span className='font-semibold opacity-70 2xl:text-[20px] lg:text-[17px] md:text-[25px] text-[18px]'>
									{obj.name}
								</span>
								<div
									className={`flex justify-center flex-row w-[${width}px] h-[${height}px]`}
								>
									<Image
										className='md:w-auto md:h-auto w-[75%] h-[75%]'
										src={obj.icon.src}
										width={width}
										height={height}
										alt=''
									/>
								</div>
							</Link>
						</li>
					);
				})}
			</ul>
			<div
				onClick={hideNav}
				className='lg:hidden block fixed top-0 left-0 z-[-1] bg-black opacity-70 w-full h-full'
			></div>
		</nav>
	);
};

export default HeaderNav;
