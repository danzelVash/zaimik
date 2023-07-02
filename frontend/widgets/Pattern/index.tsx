'use client';

import Image from 'next/image';
import { useState } from 'react';

const Pattern: React.FC = () => {
	const [loaded, setLoaded] = useState<Boolean>(false);
	return (
		<div className={`${loaded ? 'opacity-70' : 'opacity-0'} transition-opacity duration-300 fixed top-0 left-0 z-[-1] w-full h-full`}>
			<Image className='object-cover lock-padding' src="/pattern.png" fill alt='' onLoad={() => setLoaded(true)} />
		</div>
	)
}

export default Pattern;