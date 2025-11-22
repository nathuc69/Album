export async function loadData() {
    try {
        console.log('Starting fetch...');
        const IMG_RESPONSE = await fetch('http://localhost:8080/photos');
        
        if (!IMG_RESPONSE.ok) {
            console.error('HTTP error:', IMG_RESPONSE.status, IMG_RESPONSE.statusText);
            return;
        }
        
        const IMG_DATA = await IMG_RESPONSE.json();     
        //console.log('Data received:', IMG_DATA);
        return IMG_DATA;
    } catch (error) {
        console.error('Fetch error:', error.message || error);
    }
}

export async function DisplayData() {
    const CONATINER = document.getElementById('photos-container');
    const BUTTON  = document.querySelector('.more-photos');

    const DATA_IMAGES = await loadData();
    if (!DATA_IMAGES) return;

    DATA_IMAGES.forEach(image => {
        // Support multiple possible JSON key casings
        const imgPath = image.Path || image.path || image.Pathname || image.filename || image.Filename || '';
        const captionText = image.DateTaken || image.dateTaken || image.Date || image.Filename || image.filename || 'Souvenir';

        const card = document.createElement('figure');
        card.className = 'photo-card';

        const imgElement = document.createElement('img');
        imgElement.src = "../../" + imgPath; // relative path from templates -> root photos folder
        imgElement.alt = captionText || 'Photo';

        const caption = document.createElement('figcaption');
        caption.className = 'caption-handwritten';
        caption.textContent = formatCaption(captionText);

        const meta = document.createElement('div');
        meta.className = 'caption-meta';
        meta.textContent = (image.CameraMake ? image.CameraMake + ' • ' : '') + (image.CameraModel || '');

        card.appendChild(imgElement);
        card.appendChild(caption);
        card.appendChild(meta);
        CONATINER.appendChild(card);
    });
}

function formatCaption(raw) {
    // Try to format date-ish captions into a friendly form, otherwise keep provided string
    if (!raw) return 'Souvenir';
    // If it's an ISO-ish date, keep only YYYY-MM-DD or a shorter readable form
    const isoMatch = raw.match(/(\d{4}-\d{2}-\d{2})/);
    if (isoMatch) return isoMatch[1];
    return raw;
}

