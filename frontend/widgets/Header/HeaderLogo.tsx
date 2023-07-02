import Image from 'next/image';
import Link from 'next/link';
import React from 'react';

const HeaderLogo: React.FC = () => (
	<div className='lg:w-auto w-[70px] relative z-[2]'>
		<Link href='/'>
			<Image src='/logo.svg' width={100} height={92} alt='' priority />
		</Link>
	</div>
);

export default React.memo(HeaderLogo);
