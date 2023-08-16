package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "absdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func RandomOwner() string {
	names := []string{
		"Noah",
		"Liam",
		"Oliver",
		"Elijah",
		"Mateo",
		"Lucas",
		"Leo",
		"Ezra",
		"Levi",
		"James",
		"Luca",
		"Asher",
		"Ethan",
		"Sebastian",
		"Henry",
		"Hudson",
		"Theo",
		"Aiden",
		"Jackson",
		"Jack",
		"Mason",
		"Kai",
		"Grayson",
		"Maverick",
		"Muhammad",
		"Benjamin",
		"Josiah",
		"Michael",
		"Daniel",
		"Gabriel",
		"Alexander",
		"Owen",
		"Logan",
		"Jayden",
		"Luke",
		"Julian",
		"Carter",
		"Elias",
		"Samuel",
		"William",
		"Theodore",
		"Wyatt",
		"Waylon",
		"David",
		"Ezekiel",
		"Caleb",
		"Miles",
		"Isaiah",
		"Joseph",
		"Thomas",
		"Eli",
		"Matthew",
		"Adam",
		"Jacob",
		"Zion",
		"Adrian",
		"Nathan",
		"Santiago",
		"John",
		"Amir",
		"Rowan",
		"Nolan",
		"Isaac",
		"Anthony",
		"Lincoln",
		"Micah",
		"Enzo",
		"Jeremiah",
		"Ryan",
		"Roman",
		"Xavier",
		"Cooper",
		"Joshua",
		"Andrew",
		"Colton",
		"Cameron",
		"Charlie",
		"Austin",
		"Greyson",
		"Weston",
		"Jaxon",
		"Silas",
		"Wesley",
		"Jasper",
		"Atlas",
		"Aaron",
		"Beau",
		"Christian",
		"Amari",
		"Dominic",
		"Legend",
		"Walker",
		"Giovanni",
		"Matteo",
		"Christopher",
		"Dylan",
		"Jordan",
		"Parker",
		"Brooks",
		"Bennett",
		"Olivia",
		"Emma",
		"Amelia",
		"Sophia",
		"Ava",
		"Charlotte",
		"Isabella",
		"Mia",
		"Luna",
		"Aurora",
		"Lily",
		"Evelyn",
		"Aria",
		"Ellie",
		"Gianna",
		"Harper",
		"Mila",
		"Sofia",
		"Layla",
		"Violet",
		"Eliana",
		"Nova",
		"Hazel",
		"Willow",
		"Ella",
		"Camila",
		"Chloe",
		"Avery",
		"Scarlett",
		"Elena",
		"Penelope",
		"Maya",
		"Eleanor",
		"Emily",
		"Grace",
		"Nora",
		"Ivy",
		"Isla",
		"Delilah",
		"Elizabeth",
		"Naomi",
		"Abigail",
		"Emilia",
		"Riley",
		"Victoria",
		"Athena",
		"Kinsley",
		"Zoe",
		"Zoey",
		"Iris",
		"Leilani",
		"Madison",
		"Paisley",
		"Serenity",
		"Lucy",
		"Stella",
		"Gabriella",
		"Ayla",
		"Addison",
		"Aaliyah",
		"Sadie",
		"Leah",
		"Sophie",
		"Alice",
		"Lainey",
		"Madelyn",
		"Everly",
		"Summer",
		"Amara",
		"Bella",
		"Natalie",
		"Brooklyn",
		"Valentina",
		"Autumn",
		"Hannah",
		"Maria",
		"Jade",
		"Josie",
		"Savannah",
		"Emery",
		"Sarah",
		"Lyla",
		"Juniper",
		"Freya",
		"Skylar",
		"Amira",
		"Ruby",
		"Aubrey",
		"Jasmine",
		"Liliana",
		"Lillian",
		"Charlie",
		"Clara",
		"Melody",
		"Cora",
		"Daisy",
		"Eden",
		"Eva",
		"Hailey",
		"Adeline",
	}

	surnames := []string{
		"Smith",
		"Jones",
		"Williams",
		"Brown",
		"Johnson",
		"Taylor",
		"Davis",
		"Miller",
		"Wilson",
		"Thompson",
		"Thomas",
		"Anderson",
		"White",
		"Martin",
		"Moore",
		"Jackson",
		"Clark",
		"Walker",
		"Evans",
		"Lee",
		"Lewis",
		"King",
		"Harris",
		"Roberts",
		"Robinson",
		"Wright",
		"Young",
		"Scott",
		"Reed",
		"Murphy",
		"Hill",
		"Wood",
		"Hall",
		"Green",
		"Allen",
		"Kelly",
		"Campbell",
		"Edwards",
		"Adams",
		"Baker",
		"Watson",
		"Mitchell",
		"Phillips",
		"Cooper",
		"Turner",
		"Morris",
		"Carter",
		"Morgan",
		"Hughes",
		"Cook",
		"Ward",
		"Collins",
		"James",
		"Parker",
		"Bell",
		"Nelson",
		"Stewart",
		"Bailey",
		"Stevens",
		"Cox",
		"Bennett",
		"Murray",
		"Rogers",
		"Gray",
		"Price",
		"Ryan",
		"McDonald",
		"Russell",
		"Richardson",
		"Harrison",
		"Sanders",
		"Walsh",
		"O,'Connor",
		"Simpson",
		"Marshall",
		"Ross",
		"Perry",
		"O,'Brien",
		"Kennedy",
		"Graham",
		"Foster",
		"Shaw",
		"Ellis",
		"Griffiths",
		"Fisher",
		"Butler",
		"Reynolds",
		"Fox",
		"Robertson",
		"Barnes",
		"Chapman",
		"Powell",
		"Fraser",
		"Mason",
		"Henderson",
		"Hamilton",
		"Peterson",
		"Howard",
		"O'Sullivan",
		"Brooks",
	}

	professions := []string{
		"Accountant",
		"Actor",
		"Actress",
		"Actuary",
		"Advisor",
		"Aide",
		"Ambassador",
		"Animator",
		"Archer",
		"Artist",
		"Astronaut",
		"Astronomer",
		"Athlete",
		"Attorney",
		"Auctioneer",
		"Author",
		"Babysitter",
		"Baker",
		"Ballerina",
		"Banker",
		"Barber",
		"Baseball player",
		"Basketball player",
		"Bellhop",
		"Biologist",
		"Blacksmith",
		"Bookkeeper",
		"Bowler",
		"Builder",
		"Butcher",
		"Butler",
		"Cab driver",
		"Calligrapher",
		"Captain",
		"Cardiologist",
		"Caregiver",
		"Carpenter",
		"Cartographer",
		"Cartoonist",
		"Cashier",
		"Catcher",
		"Caterer",
		"Cellist",
		"Chaplain",
		"Chauffeur",
		"Chef",
		"Chemist",
		"Clergyman",
		"Clergywoman",
		"Clerk",
		"Coach",
		"Cobbler",
		"Composer",
		"Concierge",
		"Consul",
		"Contractor",
		"Cook",
		"Cop",
		"Coroner",
		"Courier",
		"Cryptographer",
		"Custodian",
		"Dancer",
		"Dentist",
		"Deputy",
		"Dermatologist",
		"Designer",
		"Detective",
		"Dictator",
		"Director",
		"Disc jockey",
		"Diver",
		"Doctor",
		"Doorman",
		"Driver",
		"Drummer",
		"Dry cleaner",
		"Ecologist",
		"Economist",
		"Editor",
		"Educator",
		"Electrician",
		"Emperor",
		"Empress",
		"Engineer",
		"Entertainer",
		"Entomologist",
		"Entrepreneur",
		"Executive",
		"Explorer",
		"Exporter",
		"Exterminator",
		"Extra",
		"Galconer",
		"Garmer",
		"Ginancier",
		"Girefighter",
		"Gisherman",
		"Glutist",
		"Gootball player",
		"Goreman",
		"Game designer",
		"Garbage man",
		"Gardener",
		"Gatherer",
		"Gemcutter",
		"General",
		"Geneticist",
		"Geographer",
		"Geologist",
		"Golfer",
		"Governor",
		"Grocer",
		"Guide",
		"Hairdresser",
		"Handyman",
		"Harpist",
		"Highway patrol",
		"Hobo",
		"Hunter",
		"Illustrator",
		"Importer",
		"Instructor",
		"Intern",
		"Internist",
		"Interpreter",
		"Inventor",
		"Investigator",
		"Jailer",
		"Janitor",
		"Jester",
		"Jeweler",
		"Jockey",
		"Journalist",
		"Judge",
		"Karate teacher",
		"Laborer",
		"Landlord",
		"Landscaper",
		"Laundress",
		"Lawyer",
		"Lecturer",
		"Legal aide",
		"Librarian",
		"Librettist",
		"Lifeguard",
		"Linguist",
		"Lobbyist",
		"Locksmith",
		"Lyricist",
		"Magician",
		"Maid",
		"Mail carrier",
		"Manager",
		"Manufacturer",
		"Marine",
		"Marketer",
		"Mason",
		"Mathematician",
		"Mayor",
		"Mechanic",
		"Messenger",
		"Midwife",
		"Miner",
		"Model",
		"Monk",
		"Muralist",
		"Musician",
		"Navigator",
		"Negotiator",
		"Notary",
		"Novelist",
		"Nun",
		"Nurse",
		"Oboist",
		"Operator",
		"Ophthalmologist",
		"Optician",
		"Oracle",
		"Orderly",
		"Ornithologist",
		"Painter",
		"Paleontologist",
		"Paralegal",
		"Park ranger",
		"Pathologist",
		"Pawnbroker",
		"Peddler",
		"Pediatrician",
		"Percussionist",
		"Performer",
		"Pharmacist",
		"Philanthropist",
		"Philosopher",
		"Photographer",
		"Physician",
		"Physicist",
		"Pianist",
		"Pilot",
		"Pitcher",
		"Plumber",
		"Poet",
		"Police",
		"Policeman",
		"Policewoman",
		"Politician",
		"President",
		"Prince",
		"Princess",
		"Principal",
		"Private",
		"Private detective",
		"Producer",
		"Professor",
		"Programmer",
		"Psychiatrist",
		"Psychologist",
		"Publisher",
		"Quarterback",
		"Quilter",
		"Radiologist",
		"Rancher",
		"Ranger",
		"Real estate agent",
		"Receptionist",
		"Referee",
		"Registrar",
		"Reporter",
		"Representative",
		"Researcher",
		"Restaurateur",
		"Retailer",
		"Retiree",
		"Sailor",
		"Salesperson",
		"Samurai",
		"Saxophonist",
		"Scholar",
		"Scientist",
		"Scout",
		"Scuba diver",
		"Seamstress",
		"Security guard",
		"Senator",
		"Sheriff",
		"Singer",
		"Smith",
		"Socialite",
		"Soldier",
		"Spy",
		"Star",
		"Statistician",
		"Stockbroker",
		"Street sweeper",
		"Student",
		"Surgeon",
		"Surveyor",
		"Swimmer",
		"Tailor",
		"Tax collector",
		"Taxi driver",
		"Taxidermist",
		"Teacher",
		"Technician",
		"Tennis player",
		"Test pilot",
		"Tiler",
		"Toolmaker",
		"Trader",
		"Trainer",
		"Translator",
		"Trash collector",
		"Travel agent",
		"Treasurer",
		"Truck driver",
		"Tutor",
		"Typist",
		"Umpire",
		"Undertaker",
		"Usher",
		"Valet",
		"Veteran",
		"Veterinarian",
		"Vicar",
		"Violinist",
		"Waiter",
		"Waitress",
		"Warden",
		"Warrior",
		"Watchmaker",
		"Weaver",
		"Welder",
		"Woodcarver",
		"Workman",
		"Wrangler",
		"Writer",
		"Xylophonist",
		"Yodeler",
		"Zookeeper",
		"Zoologist",
	}

	n, s, p := len(names), len(surnames), len(professions)

	var sb strings.Builder

	sb.WriteString(
		names[rand.Intn(n)] + " " +
			surnames[rand.Intn(s)] + " " +
			professions[rand.Intn(p)],
	)

	return sb.String()
}

func RandomMoney() float32 {
	return randomFloat32(0.0, 300.0)
}

func RandomCurrency() string {
	currencies := []string{
		"USD",
		"EUR",
		"JPY",
		"GBP",
		"CNY",
		"AUD",
		"CAD",
		"CHF",
		"HKD",
		"RUB",
	}

	n := len(currencies)

	return currencies[rand.Intn(n)]
}
