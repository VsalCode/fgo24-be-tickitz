INSERT INTO genres (name) VALUES
('Action'), -- id 1, dst
('Adventure'),
('Animation'),
('Comedy'),
('Crime'),
('Documentary'),
('Drama'),
('Family'),
('Fantasy'),
('History'),
('Horror'),
('Music'),
('Mystery'),
('Romance'),
('Science Fiction'),
('TV Movie'),
('Thriller'),
('War'),
('Western');

INSERT INTO directors (name) VALUES
('Christopher Nolan'), -- id 1, dst
('Steven Spielberg'),
('Quentin Tarantino'),
('Martin Scorsese'),
('James Cameron'),
('Ridley Scott'),
('Alfred Hitchcock'),
('George Lucas'),
('Francis Ford Coppola'),
('Woody Allen');

INSERT INTO casts (name) VALUES
('Leonardo DiCaprio'), -- id 1, dst
('Joseph Gordon-Levitt'),
('Ellen Page'),
('Christian Bale'),
('Heath Ledger'),
('Aaron Eckhart'),
('Sam Neill'),
('Laura Dern'),
('Jeff Goldblum'),
('Liam Neeson'),
('Ben Kingsley'),
('Ralph Fiennes'),
('John Travolta'),
('Uma Thurman'),
('Samuel L. Jackson'),
('Brad Pitt'),
('Christoph Waltz'),
('Mélanie Laurent'),
('Jonah Hill'),
('Margot Robbie'),
('Kate Winslet'),
('Sam Worthington'),
('Zoe Saldana'),
('Sigourney Weaver'),
('Russell Crowe'),
('Joaquin Phoenix'),
('Connie Nielsen'),
('Tom Skerritt'),
('John Hurt'),
('Anthony Perkins'),
('Janet Leigh'),
('Vera Miles'),
('James Stewart'),
('Grace Kelly'),
('Thelma Ritter'),
('Mark Hamill'),
('Harrison Ford'),
('Carrie Fisher'),
('Ewan McGregor'),
('Natalie Portman'),
('Liam Neeson'),
('Marlon Brando'),
('Al Pacino'),
('James Caan'),
('Martin Sheen'),
('Robert Duvall'),
('Woody Allen'),
('Diane Keaton'),
('Owen Wilson'),
('Rachel McAdams'),
('Marion Cotillard');

SELECT u.id, p.fullname, u.email, u.password, p.phone FROM users u 
JOIN profiles p ON u.profile_id = p.id 
WHERE u.id = 2 AND roles = 'user';

INSERT INTO payment_method (name)
VALUES ('dana'),
('google pay'),
('bca'),
('ovo'),
('paypal'),
('gopay'),
('visa'),
('bri');

INSERT INTO movies (admin_id, backdrop_path, overview, popularity, poster_path, release_date, runtime, title, vote_average) VALUES
(1, 'https://image.tmdb.org/t/p/w1280/nKyBbFSooRPTJVqjrDteD1lF733.jpg', 'After a family tragedy, kung fu prodigy Li Fong is uprooted from his home in Beijing and forced to move to New York City with his mother. When a new friend needs his help, Li enters a karate competition – but his skills alone aren''t enough. Li''s kung fu teacher Mr. Han enlists original Karate Kid Daniel LaRusso for help, and Li learns a new way to fight, merging their two styles into one for the ultimate martial arts showdown.', 612.5718, 'https://image.tmdb.org/t/p/w500/AEgggzRr1vZCLY86MAp93li43z.jpg', '2025-05-08', 123, 'Karate Kid: Legends', 7.236),
(1, 'https://image.tmdb.org/t/p/w1280/x58Gk2ZGU5AEBp25MQe2nhZhd5z.jpg', 'Andy and her team of immortal warriors fight with renewed purpose as they face a powerful new foe threatening their mission to protect humanity.', 470.299, 'https://image.tmdb.org/t/p/w500/wqfu3bPLJaEWJVk3QOm0rKhxf1A.jpg', '2025-07-01', 118, 'The Old Guard 2', 6.12),
(1, 'https://image.tmdb.org/t/p/w1280/sItIskd5xpiE64bBWYwZintkGf3.jpg', 'Taking place during the events of John Wick: Chapter 3 – Parabellum, Eve Macarro begins her training in the assassin traditions of the Ruska Roma.', 480.9303, 'https://image.tmdb.org/t/p/w500/2VUmvqsHb6cEtdfscEA6fqqVzLg.jpg', '2025-06-04', 111, 'Ballerina', 7.437),
(1, 'https://image.tmdb.org/t/p/w1280/xABhldZaMb6wfCH5oigV333OYnb.jpg', 'The UK Prime Minister and US President have a public rivalry that risks their countries'' alliance. But when they become targets of a powerful enemy, they''re forced to rely on each other as they go on a wild, multinational run. Allied with Noel, a brilliant MI6 agent, they must find a way to thwart a conspiracy that threatens the free world.', 403.4917, 'https://image.tmdb.org/t/p/w500/lVgE5oLzf7ABmzyASEVcjYyHI41.jpg', '2025-07-02', 115, 'Heads of State', 6.955),
(1, 'https://image.tmdb.org/t/p/w1280/962KXsr09uK8wrmUg9TjzmE7c4e.jpg', 'Big rig ice road driver Mike McCann travels to Nepal to scatter his late brother’s ashes on Mt. Everest. While on a packed tour bus traversing the deadly 12,000 ft. terrain of the infamous Road to the Sky, McCann and his mountain guide encounter a group of mercenaries and must fight to save themselves, the busload of innocent travelers, and the local villagers’ homeland.', 444.7102, 'https://image.tmdb.org/t/p/w500/2vHQBX5L4yoExTa55KmGIg2Q5s3.jpg', '2025-06-27', 108, 'Ice Road: Vengeance', 6.972),
(1, 'https://image.tmdb.org/t/p/w1280/g62G6aBcAcJv3ClCKmJgmHarHvq.jpg', 'Superman, a journalist in Metropolis, embarks on a journey to reconcile his Kryptonian heritage with his human upbringing as Clark Kent.', 352.4649, 'https://image.tmdb.org/t/p/w500/ombsmhYUqR4qqOLOxAyr5V8hbyv.jpg', '2025-07-09', 125, 'Superman', 7.3),
(1, 'https://image.tmdb.org/t/p/w1280/peAzdLKtT6VDWIfPQO9LJD1NCG4.jpg', 'Five years after the events of Jurassic World Dominion, covert operations expert Zora Bennett is contracted to lead a skilled team on a top-secret mission to secure genetic material from the world''s three most massive dinosaurs...', 334.7735, 'https://image.tmdb.org/t/p/w500/q0fGCmjLu42MPlSO9OYWpI5w86I.jpg', '2025-07-01', 128, 'Jurassic World Rebirth', 6.4),
(1, 'https://image.tmdb.org/t/p/w1280/9A0wQG38VdEu3DYh8HzXKXKhA6g.jpg', 'Dora, Diego, and their new friends trek through the perilous dangers of the Amazonian jungle in search of the ancient and powerful treasure of Sol Dorado...', 231.5769, 'https://image.tmdb.org/t/p/w500/r3d6u2n7iPoWNsSWwlJJWrDblOH.jpg', '2025-07-02', 96, 'Dora and the Search for Sol Dorado', 7.038),
(1, 'https://image.tmdb.org/t/p/w1280/7Zx3wDG5bBtcfk8lcnCWDOLM4Y4.jpg', 'The wildly funny and touching story of a lonely Hawaiian girl and the fugitive alien who helps to mend her broken family.', 210.4407, 'https://image.tmdb.org/t/p/w500/tUae3mefrDVTgm5mRzqWnZK6fOP.jpg', '2025-05-17', 95, 'Lilo & Stitch', 7.141),
(1, 'https://image.tmdb.org/t/p/w1280/tdMbbFhqyEqOK1QzNTvJjHWKbZX.jpg', 'In a fight for survival against a horrifying army of zombies, a former Muay Thai fighter must use skill, speed and grit to save his girlfriend.', 233.4119, 'https://image.tmdb.org/t/p/w500/7ZRXha6VZEGP3C8Kab7pPwSMzDZ.jpg', '2025-07-09', 102, 'Ziam', 6.885);



INSERT INTO movies (admin_id, backdrop_path, overview, popularity, poster_path, release_date, runtime, title, vote_average) VALUES
(1, 'https://image.tmdb.org/t/p/w1280/nKyBbFSooRPTJVqjrDteD1lF733.jpg', 'After a family tragedy, kung fu prodigy Li Fong is uprooted from his home in Beijing and forced to move to New York City with his mother. When a new friend needs his help, Li enters a karate competition – but his skills alone aren''t enough. Li''s kung fu teacher Mr. Han enlists original Karate Kid Daniel LaRusso for help, and Li learns a new way to fight, merging their two styles into one for the ultimate martial arts showdown.', 612.5718, 'https://image.tmdb.org/t/p/w500/AEgggzRr1vZCLY86MAp93li43z.jpg', '2025-05-08', 130, 'Karate Kid: Legends', 7.236),
(1, 'https://image.tmdb.org/t/p/w1280/sItIskd5xpiE64bBWYwZintkGf3.jpg', 'Taking place during the events of John Wick: Chapter 3 – Parabellum, Eve Macarro begins her training in the assassin traditions of the Ruska Roma.', 480.9303, 'https://image.tmdb.org/t/p/w500/2VUmvqsHb6cEtdfscEA6fqqVzLg.jpg', '2025-06-04', 115, 'Ballerina', 7.432),
(1, 'https://image.tmdb.org/t/p/w1280/5esDYWV0NoFwqPa1iC0g9akqZo9.jpg', 'Following the death of their father, a brother and sister are introduced to their new sibling by their foster mother, only to learn that she has a terrifying secret.', 178.5552, 'https://image.tmdb.org/t/p/w500/tObSf1VzzHt9xB0csanFtb3DRjf.jpg', '2025-05-28', 100, 'Bring Her Back', 7.369),
(1, 'https://image.tmdb.org/t/p/w1280/yAqL0makiGke5yYiVWpmBDSKIVP.jpg', 'Two priests, one in crisis with his faith and the other confronting a turbulent past, must overcome their differences to perform a risky exorcism.', 137.9185, 'https://image.tmdb.org/t/p/w500/ktqPs5QyuF8SpKnipvVHb3fwD8d.jpg', '2025-04-18', 102, 'The Ritual', 6.011),
(1, 'https://image.tmdb.org/t/p/w1280/kyBOGOBUMdGWOhzECuosPSzoMpi.jpg', 'After the underlying tech for M3GAN is stolen and misused by a powerful defense contractor to create a military-grade weapon known as Amelia, M3GAN''s creator Gemma realizes that the only option is to resurrect M3GAN and give her a few upgrades, making her faster, stronger, and more lethal.', 71.2551, 'https://image.tmdb.org/t/p/w500/4a63rQqIDTrYNdcnTXdPsQyxVLo.jpg', '2025-06-25', 110, 'M3GAN 2.0', 7.205);

INSERT INTO movies (admin_id, backdrop_path, overview, popularity, poster_path, release_date, runtime, title, vote_average) VALUES
(1, 'https://image.tmdb.org/t/p/w1280/k3oqTinhaXFKxlY09oYxnSks9R0.jpg', 'Jérémie returns to his hometown for the funeral of his former boss, the village baker. He decides to stay for a few days with Martine, the man''s widow. A mysterious disappearance, a threatening neighbor and a priest with strange intentions make Jérémie''s short stay in the village take an unexpected turn.', 46.8443, 'https://image.tmdb.org/t/p/w500/i1eFKvcUNdyNd53YeEOK2vNkBsL.jpg', '2024-10-16', 108, 'Misericordia', 6.3),
(1, 'https://image.tmdb.org/t/p/w1280/unthV1mq9llhEinIMPcCUImFodt.jpg', 'Against the vibrant backdrop of a 1960s-inspired, retro-futuristic world, Marvel''s First Family is forced to balance their roles as heroes with the strength of their family bond, while defending Earth from a ravenous space god called Galactus and his enigmatic Herald, Silver Surfer.', 46.4319, 'https://image.tmdb.org/t/p/w500/x26MtUlwtWD26d0G0FXcppxCJio.jpg', '2025-07-23', 140, 'The Fantastic Four: First Steps', 0),
(1, 'https://image.tmdb.org/t/p/w1280/w3RDV3pSpxN0C2DZ4Xpw4o5LWpI.jpg', 'Wealthy businessman Zsa-zsa Korda appoints his only daughter, a nun, as sole heir to his estate. As Korda embarks on a new enterprise, they soon become the target of scheming tycoons, foreign terrorists, and determined assassins.', 42.6036, 'https://image.tmdb.org/t/p/w500/jtEqpy0K1iVuCebRwWqT6BZ6jLN.jpg', '2025-05-23', 125, 'The Phoenician Scheme', 6.801),
(1, 'https://image.tmdb.org/t/p/w1280/icFWIk1KfkWLZnugZAJEDauNZ94.jpg', 'One year after her sister Melanie mysteriously disappeared, Clover and her friends head into the remote valley where she vanished in search of answers. Exploring an abandoned visitor center, they find themselves stalked by a masked killer and horrifically murdered one by one...only to wake up and find themselves back at the beginning of the same evening.', 37.0115, 'https://image.tmdb.org/t/p/w500/juA4IWO52Fecx8lhAsxmDgy3M3.jpg', '2025-04-23', 98, 'Until Dawn', 6.6),
(1, 'https://image.tmdb.org/t/p/w1280/21lM9hfYNDX9RnWzp4aq0iQTEbq.jpg', 'Mafuyu doubts his place in the band, but a familiar face could be the spark that brings him back to the stage.', 35.2115, 'https://image.tmdb.org/t/p/w500/2VLnT5JtJ77ErG6GD3c1J628VLp.jpg', '2024-09-20', 94, 'given the Movie: To the Sea', 8.333);


INSERT INTO movie_genres (movie_id, genre_id) VALUES
-- Action (1)
(1,1), (2,1), (3,1), (4,1), (5,1), (6,1), (7,1), (13,1), (15,1),
-- Adventure (2)
(2,2), (5,2), (6,2), (7,2), (8,2), (15,2), (16,2), (18,2),
-- Animation (3)
(9,3), 
-- Comedy (4)
(4,4),
-- Drama (7)
(1,7), (14,7), (18,7),
-- Family (8)
(1,8), (8,8), (9,8), (18,8),
-- Fantasy (9)
(2,9), (17,9),
-- Horror (11)
(12,11), (14,11), (19,11),
-- Mystery (13)
(12,13), (14,13), (19,13),
-- Science Fiction (15)
(6,15), (7,15), (15,15),
-- Thriller (17)
(3,17), (4,17), (5,17), (11,17), (16,17);


INSERT INTO movie_directors (movie_id, director_id) VALUES
(1,2),   -- Steven Spielberg
(2,5),   -- James Cameron
(3,3),   -- Quentin Tarantino
(4,4),   -- Martin Scorsese
(5,6),   -- Ridley Scott
(6,1),   -- Christopher Nolan
(7,2),   -- Steven Spielberg
(8,10),  -- Woody Allen
(9,10),  -- Woody Allen
(10,1),  -- Christopher Nolan
(11,7),  -- Alfred Hitchcock
(12,9),  -- Francis Ford Coppola
(13,8),  -- George Lucas
(14,9),  -- Francis Ford Coppola
(15,2),  -- Steven Spielberg
(16,6),  -- Ridley Scott
(17,1),  -- Christopher Nolan
(18,5),  -- James Cameron
(19,7),  -- Alfred Hitchcock
(20,5);  -- James Cameron

INSERT INTO movie_casts (movie_id, cast_id) VALUES
-- Film 1: Karate Kid
(1,1), (1,4), (1,10), 
-- Film 2: The Old Guard 2
(2,16), (2,20), (2,21), 
-- Film 3: Ballerina
(3,14), (3,20), (3,50), 
-- Film 4: Heads of State
(4,1), (4,16), (4,19), 
-- Film 5: Ice Road: Vengeance
(5,10), (5,22), (5,25), 
-- Film 6: Superman
(6,4), (6,6), (6,21), 
-- Film 7: Jurassic World Rebirth
(7,7), (7,8), (7,9), 
-- Film 8: Dora
(8,31), (8,32), (8,45), 
-- Film 9: Lilo & Stitch
(9,46), (9,47), (9,48), 
-- Film 10: Ziam
(10,1), (10,2), (10,3), 
-- Film 11: Bring Her Back
(11,28), (11,29), (11,30), 
-- Film 12: The Ritual
(12,24), (12,26), (12,27), 
-- Film 13: MAGAN 2.0
(13,36), (13,37), (13,38), 
-- Film 14: Misericordia
(14,40), (14,41), (14,42), 
-- Film 15: Fantastic Four
(15,23), (15,39), (15,49), 
-- Film 16: Phoenician Scheme
(16,10), (16,11), (16,12), 
-- Film 17: Until Dawn
(17,30), (17,33), (17,34), 
-- Film 18: To the Sea
(18,21), (18,22), (18,45);